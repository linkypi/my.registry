package tcp

import (
	common "github.com/linkypi/hiraeth.registry/common"
	"github.com/linkypi/hiraeth.registry/server/api/handler"
	"github.com/linkypi/hiraeth.registry/server/cluster"
	"github.com/linkypi/hiraeth.registry/server/config"
	"github.com/linkypi/hiraeth.registry/server/log"
	cpb "github.com/linkypi/hiraeth.registry/server/proto"
	"github.com/linkypi/hiraeth.registry/server/slot"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pkg/logging"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"strconv"
	"time"
)

type Server struct {
	*gnet.EventServer
	log         *logrus.Logger
	addr        string
	cluster     *cluster.Cluster
	slotManager *slot.Manager
	startUpMode config.StartUpMode
	netManager  *NetManager
	shutDownCh  chan struct{}

	Ch chan RequestWrapper

	heartbeatSec   int
	workerPool     *goroutine.Pool
	requestHandler handler.RequestHandlerFactory
}

func NewClientTcpServer(addr string, cluster *cluster.Cluster, startupMode config.StartUpMode,
	slotManager *slot.Manager, shutDownCh chan struct{}) *Server {

	netManager := NewNetManager()
	serviceImpl := handler.ServiceImpl{
		SlotManager: slotManager,
		Cluster:     cluster, Log: log.Log,
		StartUpMode: startupMode}
	handlerFactory := handler.NewHandlerFactory(&serviceImpl, log.Log)
	tcpServer := Server{
		log:            log.Log,
		addr:           addr,
		Ch:             make(chan RequestWrapper, 1000),
		cluster:        cluster,
		slotManager:    slotManager,
		startUpMode:    startupMode,
		netManager:     netManager,
		heartbeatSec:   3,
		workerPool:     goroutine.Default(),
		shutDownCh:     shutDownCh,
		requestHandler: handlerFactory,
	}
	serviceImpl.OnSubEvent = tcpServer.publishServiceChanged

	return &tcpServer
}

func (s *Server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Log.Infof("client tcp server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}
func (s *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	s.log.Infof("connected to client: %s", c.RemoteAddr().String())
	s.netManager.AddConn(c)
	return
}

func (s *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	s.log.Infof("client connection is closed, remote addr: %s, error: %s", c.RemoteAddr().String(), err.Error())
	s.netManager.RemoveConn(c.RemoteAddr().String())
	return
}

func (s *Server) Tick() (delay time.Duration, action gnet.Action) {
	s.netManager.SendHeartbeat()
	delay = time.Duration(s.heartbeatSec) * time.Millisecond
	return
}

func (s *Server) React(buffer []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	data := append([]byte{}, buffer...)
	msgType := data[0]
	data = data[1:]
	s.Ch <- RequestWrapper{Data: data, Conn: c, MsgType: common.MessageType(msgType)}
	return
}

func (s *Server) publishServiceChanged(connIds []string, serviceName string, instances []common.ServiceInstance) {
	if s.netManager == nil {
		return
	}

	// check cluster state
	if s.startUpMode == config.Cluster && s.cluster.State != cluster.Active {
		s.log.Warnf("failed to publish service changed: %s, cluster %d is not active: %s",
			serviceName, s.cluster.ClusterId, s.cluster.State.String())
		return
	}

	request := cpb.ServiceChangedRequest{ServiceName: serviceName}

	for _, connId := range connIds {
		conn := s.netManager.GetConn(connId)
		if conn == nil {
			s.log.Warnf("failed to publish service changed: %s, conn [%s] not exist", serviceName, connId)
			continue
		}
		payload, err := proto.Marshal(&request)
		if err != nil {
			s.log.Errorf("failed to publish service changed: %s, %s", serviceName, err.Error())
			return
		}
		newRequest := common.NewRequest(common.PublishServiceChanged, payload)
		bytes, err := newRequest.ToBytes()
		if err != nil {
			s.log.Errorf("failed to publish service changed: %s, encode occur error: %s", serviceName, err.Error())
			return
		}
		err = conn.AsyncWrite(bytes)
		if err != nil {
			s.log.Errorf("failed to publish service changed: %s, async reply occur error: %s", serviceName, err.Error())
		}
	}
}

func (s *Server) OnShutdown(server gnet.Server) {
	s.log.Errorf("client server shutting down")
	s.Shutdown()
}

func (s *Server) Start(nodeId string) {

	if !s.initSnowFlake(nodeId) {
		return
	}

	startCh := make(chan struct{})
	go s.handleRequest(startCh)
	select {
	case <-startCh:
	}

	go s.startTcpServer(s.addr)

	select {
	case <-s.shutDownCh:
		s.Shutdown()
		return
	}
}

func (s *Server) initSnowFlake(nodeId string) bool {
	machineId, err := strconv.Atoi(nodeId)
	if err != nil {
		machineId = 1
		s.log.Warnf("invalid nodeId, use default nodeId: %d", machineId)
	}

	err = common.InitSnowFlake("", int64(machineId))
	if err != nil {
		s.log.Errorf("init snowflake failed, err: %v", err)
		close(s.shutDownCh)
		return false
	}
	return true
}

func (s *Server) startTcpServer(addr string) {

	defer func() {
		if err := recover(); err != nil {
			s.log.Errorf("start tcp server panic: %v", err)
		}
	}()

	codec := gnet.NewLengthFieldBasedFrameCodec(common.EncoderConfig, common.DecoderConfig)
	err := gnet.Serve(s, addr,
		gnet.WithLogger(log.Log),
		gnet.WithLogLevel(logging.DebugLevel),
		gnet.WithCodec(codec),
		gnet.WithMulticore(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithReusePort(true),
		gnet.WithReuseAddr(true),
		gnet.WithNumEventLoop(4),
		gnet.WithTCPKeepAlive(time.Minute*5),
	)

	if err != nil {
		s.log.Errorf("start client server failed, err: %v", err)
		s.Shutdown()
	}
}

func (s *Server) Shutdown() {
	s.netManager.CloseAllConn()
	close(s.shutDownCh)
}

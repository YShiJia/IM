/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 18:48:18
 */

package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/YShiJia/IM/apps/edge/internal/logic/pushreceivemsg"
	"github.com/YShiJia/IM/apps/edge/internal/logic/recvcmdmsg"
	"github.com/YShiJia/IM/apps/edge/internal/logic/recvcommonmsg"
	"github.com/YShiJia/IM/apps/edge/internal/middleware"
	"github.com/YShiJia/IM/apps/edge/internal/svc"
	"github.com/YShiJia/IM/lib/discovery"
	"github.com/YShiJia/IM/lib/ip"
	"github.com/YShiJia/IM/lib/wait"
	"github.com/YShiJia/IM/lib/webscoket"
	"github.com/YShiJia/IM/lib/webscoket/conn"
	"github.com/YShiJia/IM/model"
	"github.com/YShiJia/IM/pbmodel/pbmessage"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/threading"
)

var defaultRecvClientMsgChanSize = 100

var edgeServerEntity *EdgeServer

func Entity() *EdgeServer {
	return edgeServerEntity
}

type EdgeServer struct {
	SvcCtx   *svc.ServiceContext
	WsServer *webscoket.WsServer

	//发送消息
	SendMsgPusher *kq.Pusher
	//消费命令消息
	RecvCmdMsgConsumer service.Service
	//消费待发送至客户端的消息
	RecvCommonMsgConsumer service.Service
	//服务组
	serviceGroup *service.ServiceGroup

	//接收消息缓冲区
	ClientMsgChan chan *pbmessage.PbMessage
	// 退出信号
	closed chan struct{}
}

func NewEdgeServer(svcCtx *svc.ServiceContext) *EdgeServer {
	var edgeserver EdgeServer

	//将conn处理流程注册
	svcCtx.WsServer.WithConnHandler(edgeserver.wsConnHandler)
	svcCtx.WsServer.MiddleWare(middleware.WsConnJwtAuthorize(svcCtx))

	RecvCmdMsgConsumer := kq.MustNewQueue(
		svcCtx.RecvCmdMsgConsumerConf,
		recvcmdmsg.NewRecvCmdMsgConsumer(context.TODO(), svcCtx),
	)
	RecvCommonMsgConsumer := kq.MustNewQueue(
		svcCtx.RecvCommonMsgConsumerConf,
		recvcommonmsg.NewRecvCommonMsgConsumer(context.TODO(), svcCtx),
	)
	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(RecvCmdMsgConsumer)
	serviceGroup.Add(RecvCommonMsgConsumer)

	edgeserver.SvcCtx = svcCtx
	edgeserver.WsServer = svcCtx.WsServer
	edgeserver.SendMsgPusher = svcCtx.SendMsgPusher
	edgeserver.RecvCmdMsgConsumer = RecvCmdMsgConsumer
	edgeserver.RecvCommonMsgConsumer = RecvCommonMsgConsumer
	edgeserver.ClientMsgChan = make(chan *pbmessage.PbMessage, defaultRecvClientMsgChanSize)
	//注册到本地
	edgeServerEntity = &edgeserver

	return &edgeserver
}

func (e *EdgeServer) Register(method, path string) error {
	return e.WsServer.Register(method, path)
}

func (e *EdgeServer) Start() error {
	// 消费者队列启动
	//threading.GoSafe(func() {
	//	e.serviceGroup.Start()
	//})
	// etcd注册心跳器
	threading.GoSafe(func() {
		e.KqHeart()
	})
	//处理客户端消息
	threading.GoSafe(func() {
		if err := e.handleClientMsg(); err != nil {
			e.SvcCtx.Errorf("handleClientMsg error: %v", err)
		}
	})
	return e.WsServer.Start()
}

func (e *EdgeServer) Stop() error {
	//threading.GoSafe(func() {
	//	e.serviceGroup.Stop()
	//})
	close(e.closed)
	e.WsServer.Stop()

	e.SvcCtx.Infof("edge server stop")
	return nil
}

func (e *EdgeServer) PushReceiveMsg(key string, msg any) error {
	pushReceiveMsgLogic := pushreceivemsg.NewPushReceiveMsgLogic(context.TODO(), e.SvcCtx)
	pushReceiveMsgLogic.PushReceiveMsgLogic(key, msg)
	return nil
}

// 这里就是处理的核心逻辑，接收来着客户端的消息，每个连接都会执行这个逻辑
func (e *EdgeServer) wsConnHandler(c conn.Conn) {
	uid, _ := c.GetValue(webscoket.WsConnectionIdentity)
	threading.GoSafe(func() {
		for {
			msg := new(pbmessage.PbMessage)
			err := c.Receive(context.TODO(), msg)
			if err != nil {
				if errors.Is(err, conn.ErrConnectionClosed) {
					return
				} else {
					e.SvcCtx.Errorf("wsConnHandler receive msg error: %v", err)
				}
			}
			e.ClientMsgChan <- msg
			e.SvcCtx.Infof("wsConnHandler receive msg: %v, from user: %s", msg, uid)
		}
	})
	value, _ := c.GetValue(webscoket.WsConnectionIdentity)
	e.SvcCtx.Infof("wsConnHandler start, conn key: %s", value)
}

// 处理来着客户端的消息,服务器来收集来着客户端的消息，并通过
func (e *EdgeServer) handleClientMsg() error {
	go func() {
		var waiter wait.Waiter = wait.NewWaiterByExponentialBackoff()
		for {
			select {
			case msg := <-e.ClientMsgChan:
				if err := ConnClientMsgCenterProcessor.ClientMsgHandle(msg); err != nil {
					e.SvcCtx.Errorf("handleClientMsg error: %v", err)
				}
				waiter.Reset()
			case <-e.closed:
				return
			default:
				// 不能让他一直空转，没有任务处理就让出CPU
				// 使用指数回退算法，计算休眠时间
				waiter.Wait()
			}
		}
	}()
	return nil
}

func (e *EdgeServer) KqHeart() error {
	// 获取本服务id地址
	addr, err := ip.GetIPv4Addr(e.SvcCtx.IPv4Prefix)
	if err != nil || len(addr) == 0 {
		return fmt.Errorf("get ip error")
	}
	//将自身服务Name作为key， 接收消息队列配置信息作为value
	edgemqinfo := &model.EdgeMQInfo{
		RecvCmdMsgConsumerConf:    e.SvcCtx.RecvCmdMsgConsumerConf,
		RecvCommonMsgConsumerConf: e.SvcCtx.RecvCommonMsgConsumerConf,
		Address:                   addr[0],
	}
	register := discovery.NewServerRegister(e.SvcCtx.Etcd.Key, e.SvcCtx.Etcd.Hosts, edgemqinfo)
	register.HeartBeat()
	return nil
}

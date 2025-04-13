package api

import (
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/YShiJia/IM/apps/edge/internal/logic"
	libWs "github.com/YShiJia/IM/lib/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

// MessageSRHandler 处理和发送来自客户端的联系
func MessageSRHandler(ctx *libWs.WebContext) {
	cwc := logic.CommunicateWithClient{
		RequestId:             ctx.Value(libWs.RequestUUID),
		WsConn:                ctx.Conn(),
		UserUid:               ctx.Value(conf.AuthIMUserUID),
		ClientSilenceTimer:    time.NewTimer(conf.Conf.ClientMaxSilenceTime),
		ClientMaxSilenceTime:  conf.Conf.ClientMaxSilenceTime,
		ErrCh:                 make(chan error),
		CloseCh:               make(chan struct{}),
		OnlineHeatBeatCloseCh: make(chan struct{}),
	}
	log.Infof("clientInfo=%+v conn[%s] start deal message", cwc, ctx.Value(libWs.RequestUUID))
	defer func() {
		close(cwc.CloseCh)
		close(cwc.CloseCh)
		close(cwc.OnlineHeatBeatCloseCh)
		if panicErr := recover(); panicErr != nil {
			if pErr, ok := panicErr.(error); ok {
				log.Errorf("conn[%s] got panic error: %+v", cwc.RequestId, errors.WithStack(pErr))
			}
		}
	}()
	go cwc.Recv()
	go cwc.HeartBeat()

	select {
	case err := <-cwc.ErrCh:
		log.Errorf("conn[%s] 发送错误，连接断开, err: %v", cwc.RequestId, err)
	case <-cwc.CloseCh:
		log.Infof("conn[%s] 正常结束，连接断开", cwc.RequestId)
	case <-cwc.ClientSilenceTimer.C:
		log.Infof("conn[%s] 客户端心跳超时，连接断开", cwc.RequestId)
	}
}

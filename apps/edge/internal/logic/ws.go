/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-04 15:51:41
 */

package logic

import (
	"context"
	"encoding/json"
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/YShiJia/IM/apps/edge/internal/dao"
	"github.com/YShiJia/IM/apps/edge/internal/model"
	libWsconn "github.com/YShiJia/IM/lib/websocket/conn"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

// CommunicateWithClient 与客户端进行通讯
type CommunicateWithClient struct {
	mutex                 sync.Mutex
	RequestId             string         // 每次请求的唯一ID
	WsConn                libWsconn.Conn // websocket连接
	UserUid               string         // 用户Uid
	ClientSilenceTimer    *time.Timer    // 客户端静默时间超时器
	ClientMaxSilenceTime  time.Duration  // 客户端最大静默时间
	ErrCh                 chan error     // 错误通道
	CloseCh               chan struct{}  // 关闭通道
	OnlineHeatBeatCloseCh chan struct{}  // 用户在线心跳机制关闭channel
}

// 重置超时时间
func (cwc *CommunicateWithClient) resetClientSilenceTimer() {
	if cwc.ClientSilenceTimer == nil {
		return
	}
	// stop失败
	if !cwc.ClientSilenceTimer.Stop() {
		// 如果计时器已经超时了, <-timer.C会阻塞, 增加default逻辑避免阻塞
		select {
		case <-cwc.ClientSilenceTimer.C:
		default:
		}
	}
	// 如果timer启动,则重新设置超时时间
	cwc.ClientSilenceTimer.Reset(cwc.ClientMaxSilenceTime)
}

func (cwc *CommunicateWithClient) Recv() {
	defer func() {
		log.Infof("conn[%s] recv goroutine exit", cwc.RequestId)
		if panicErr := recover(); panicErr != nil {
			if pErr, ok := panicErr.(error); ok {
				log.Errorf("conn[%s] got panic error: %+v", cwc.RequestId, errors.WithStack(pErr))
			}
		}
	}()

	for {
		recvType, data, err := cwc.WsConn.Receive() // 会阻塞,直到接收到消息或连接断开
		// 接收数据
		if err != nil {
			log.Errorf("conn[%s] recv error: %v", cwc.RequestId, err)
			break
		}
		if recvType == websocket.CloseNormalClosure {
			// 客户端要求连接断开
			cwc.CloseCh <- struct{}{}
			return
		}
		msg := &model.Message{}
		if err := json.Unmarshal(data, msg); err != nil {
			log.Errorf("conn[%s] recv data %s, unmarshal error: %v", cwc.RequestId, string(data), err)
		}
		// 重新开始计时
		cwc.resetClientSilenceTimer()
		log.Infof("conn[%s] recv msg: %v", cwc.RequestId, msg)
		switch msg.MessageType {
		case model.MessageTypePing: // 心跳机制
		case model.MessageTypeClose: // 关闭连接
			log.Infof("conn[%s] recv close sign msg: %v", cwc.RequestId, msg)
			cwc.CloseCh <- struct{}{}
			return
		case model.MessageTypePrivate, model.MessageTypeGroup: // 发送消息
		// TODO: 发送到kafka的总消息队列中
		default:
			log.Errorf("conn[%s] unknown message type: %v", cwc.RequestId, msg.MessageType)
		}
	}
}

// HeartBeat 将用户上下线的消息放到redis中
func (cwc *CommunicateWithClient) HeartBeat() {
	// 半个最大静默时间为一个周期，为用户在线状态续期
	keepAlivePeriod := cwc.ClientMaxSilenceTime / 2
	keepAliveTimer := time.NewTimer(keepAlivePeriod)
	for {
		select {
		case <-keepAliveTimer.C:
			if err := dao.Redis.RegisterUserUid(context.Background(), cwc.UserUid, conf.Conf.Name, cwc.ClientMaxSilenceTime); err != nil {
				log.Errorf("conn[%s] set userUid[%s] to redis error: %v", cwc.RequestId, cwc.UserUid, err)
			}
			keepAliveTimer.Reset(keepAlivePeriod)
		case <-cwc.OnlineHeatBeatCloseCh:
			return
		}
	}
}

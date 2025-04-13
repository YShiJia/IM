/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-07 19:58:54
 */

package svc

import (
	"context"
	"fmt"
	"github.com/YShiJia/IM/pbmodel/pbmessage"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgConsumerLogic struct {
	svcCtx *ServiceContext
	logx.Logger
}

func NewSendMsgConsumerLogic(svcCtx *ServiceContext) *SendMsgConsumerLogic {
	return &SendMsgConsumerLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

// TODO 这里需要加上一个落盘处理,+mq的手动commit, 可以使用协程池进行处理
func (l *SendMsgConsumerLogic) Consume(ctx context.Context, key string, value string) error {
	var msg pbmessage.PbMessage
	if err := l.svcCtx.Encoder.Decode([]byte(value), &msg); err != nil {
		//解析消息失败
		return fmt.Errorf("解析消息失败: %v", err)
	}
	//现在一般都是转发消息，没有其他需求消息，直接转发就行
	//等到以后有新的消息类型需求，再做其他处理
	nodeName, err := l.svcCtx.CsHash.GetNode(ctx, msg.To)
	if err != nil {
		return fmt.Errorf("获取节点失败: %v", err)
	}
	edgeServerInfo, ok := l.svcCtx.EdgeService.Get(nodeName)
	if !ok {
		return fmt.Errorf("节点不存在")
	}
	return edgeServerInfo.RecvCommonMsgPusher.KPush(ctx, msg.To, value)
}

/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-04 13:22:31
 */

package recvcommonmsg

import (
	"errors"
	"github.com/YShiJia/IM/pbmodel/pbmessage"
)

var CommonMsgProcessor = newCommonMsgCenterHandler(
	map[pbmessage.PbMessageType]CommonMsgHandler{
		pbmessage.PbMessageType_Transfer: transferCommonMsgHandlerEntity,
	},
)

// TODO 后期想办法提升一下代码复用率
type CommonMsgHandler interface {
	CommonMsgHandle(consumer *RecvCommonMsgConsumer, message *pbmessage.PbMessage) error
}

// 无对应消息类型的处理程序
var ErrCommonMsgHandlerNotFound = errors.New("no processor for this message type")

var _ CommonMsgHandler = (*CommonMsgCenterHandler)(nil)

// 处理中心, 使用策略模式
type CommonMsgCenterHandler struct {
	handlers map[pbmessage.PbMessageType]CommonMsgHandler
}

// 只会采取第一个map
func newCommonMsgCenterHandler(CommonMsgHandlers ...map[pbmessage.PbMessageType]CommonMsgHandler) *CommonMsgCenterHandler {
	var handlers map[pbmessage.PbMessageType]CommonMsgHandler
	if len(CommonMsgHandlers) != 0 {
		handlers = CommonMsgHandlers[0]
	} else {
		handlers = make(map[pbmessage.PbMessageType]CommonMsgHandler)
	}
	return &CommonMsgCenterHandler{
		handlers: handlers,
	}
}

func (c *CommonMsgCenterHandler) CommonMsgHandle(consumer *RecvCommonMsgConsumer, message *pbmessage.PbMessage) error {
	processor, ok := c.handlers[message.MsgType]
	if !ok {
		return ErrCommonMsgHandlerNotFound
	}
	//策略下发
	return processor.CommonMsgHandle(consumer, message)
}

// 传输消息命令处理
var _ CommonMsgHandler = (*transferCommonMsgHandler)(nil)

type transferCommonMsgHandler struct{}

var transferCommonMsgHandlerEntity = &transferCommonMsgHandler{}

// 连接不存在就不发送
func (t *transferCommonMsgHandler) CommonMsgHandle(consumer *RecvCommonMsgConsumer, message *pbmessage.PbMessage) error {
	conn := consumer.svcCtx.WsServer.GetConn(message.To)
	if conn != nil {
		conn.Send(consumer.ctx, message)
	}
	return nil
}

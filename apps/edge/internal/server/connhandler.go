/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-04 15:51:41
 */

package server

import (
	"errors"
	"fmt"
	"github.com/YShiJia/IM/pbmodel/pbmessage"
)

type ConnClientMsgHandler interface {
	ClientMsgHandle(message *pbmessage.PbMessage) error
}

var ConnClientMsgCenterProcessor = newConnClientMsgCenterHandler(
	map[pbmessage.PbMessageType]ConnClientMsgHandler{
		pbmessage.PbMessageType_Ping: pingClientMsgHandlerEntity,
		pbmessage.PbMessageType_Chat: chatClientMsgHandlerEntity,
	},
)

// 无对应消息类型的处理程序
var ErrCommonMsgHandlerNotFound = errors.New("no processor for this message type")

var _ ConnClientMsgHandler = (*ConnClientMsgCenterHandler)(nil)

// 处理中心, 使用策略模式
type ConnClientMsgCenterHandler struct {
	handlers map[pbmessage.PbMessageType]ConnClientMsgHandler
}

// 只会采取第一个map
func newConnClientMsgCenterHandler(CommonMsgHandlers ...map[pbmessage.PbMessageType]ConnClientMsgHandler) *ConnClientMsgCenterHandler {
	var handlers map[pbmessage.PbMessageType]ConnClientMsgHandler
	if len(CommonMsgHandlers) != 0 {
		handlers = CommonMsgHandlers[0]
	} else {
		handlers = make(map[pbmessage.PbMessageType]ConnClientMsgHandler)
	}
	return &ConnClientMsgCenterHandler{
		handlers: handlers,
	}
}

func (c *ConnClientMsgCenterHandler) ClientMsgHandle(message *pbmessage.PbMessage) error {
	processor, ok := c.handlers[message.MsgType]
	if !ok {
		return ErrCommonMsgHandlerNotFound
	}
	//策略下发
	return processor.ClientMsgHandle(message)
}

// 心跳检测机制处理
var _ ConnClientMsgHandler = (*pingClientMsgHandler)(nil)

type pingClientMsgHandler struct{}

var pingClientMsgHandlerEntity = &pingClientMsgHandler{}

// 刷新连接时长，什么都不用做
func (p *pingClientMsgHandler) ClientMsgHandle(message *pbmessage.PbMessage) error {
	return nil
}

// 聊天消息处理
var _ ConnClientMsgHandler = (*chatClientMsgHandler)(nil)

type chatClientMsgHandler struct{}

var chatClientMsgHandlerEntity = &chatClientMsgHandler{}

// 转发消息
func (c *chatClientMsgHandler) ClientMsgHandle(message *pbmessage.PbMessage) error {
	if err := c.checkMsg(message); err != nil {
		return err
	}
	//这一步是发送到总消息队列中，后续的处理大部分都与接受者有关，所以使用接受者id作为key
	Entity().PushReceiveMsg(message.To, message)
	return nil
}

func (c *chatClientMsgHandler) checkMsg(message *pbmessage.PbMessage) error {
	if message.MsgType != pbmessage.PbMessageType_Chat {
		return fmt.Errorf("msg type is not chat")
	}
	if message.From == "" || message.To == "" {
		return fmt.Errorf("msg's from or to is empty")
	}
	if message.Seq == 0 {
		return fmt.Errorf("msg's seq is 0")
	}
	return nil
}

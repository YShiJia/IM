/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-03 13:40:56
 */

package recvcmdmsg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YShiJia/IM/model/pbmessage"
)

var CmdMsgProcessor = newCmdMsgCenterHandler(
	map[pbmessage.PbMessageType]CmdMsgHandler{
		pbmessage.PbMessageType_UpdateConn: updateConnCmdMsgHandlerEntity,
	},
)

type CmdMsgHandler interface {
	CmdMsgHandle(consumer *RecvCmdMsgConsumer, message *pbmessage.PbMessage) error
}

// 无对应消息类型的处理程序
var ErrCmdMsgHandlerNotFound = errors.New("no processor for this message type")

var _ CmdMsgHandler = (*cmdMsgCenterHandler)(nil)

// 处理中心, 使用策略模式
type cmdMsgCenterHandler struct {
	handlers map[pbmessage.PbMessageType]CmdMsgHandler
}

// 只会采取第一个map
func newCmdMsgCenterHandler(CmdMsgHandlers ...map[pbmessage.PbMessageType]CmdMsgHandler) *cmdMsgCenterHandler {
	var handlers map[pbmessage.PbMessageType]CmdMsgHandler
	if len(CmdMsgHandlers) != 0 {
		handlers = CmdMsgHandlers[0]
	} else {
		handlers = make(map[pbmessage.PbMessageType]CmdMsgHandler)
	}
	return &cmdMsgCenterHandler{
		handlers: handlers,
	}
}

func (c *cmdMsgCenterHandler) CmdMsgHandle(consumer *RecvCmdMsgConsumer, message *pbmessage.PbMessage) error {
	processor, ok := c.handlers[message.MsgType]
	if !ok {
		return ErrCmdMsgHandlerNotFound
	}
	//策略下发
	return processor.CmdMsgHandle(consumer, message)
}

// 更新连接信息命令处理
var _ CmdMsgHandler = (*updateConnCmdMsgHandler)(nil)

type updateConnCmdMsgHandler struct{}

var updateConnCmdMsgHandlerEntity = &updateConnCmdMsgHandler{}

func (u *updateConnCmdMsgHandler) CmdMsgHandle(consumer *RecvCmdMsgConsumer, message *pbmessage.PbMessage) error {
	if message.MsgType != pbmessage.PbMessageType_UpdateConn {
		//基层代码错误
		panic("message type is not update conn")
	}
	var msg pbmessage.UpdateConnMsg
	if err := json.Unmarshal(message.Data, &msg); err != nil {
		return fmt.Errorf("Json unmarshal pbmessage.UpdateConnMsg failed: %w", err)
	}
	for _, key := range msg.Keys {
		conn := consumer.svcCtx.WsServer.GetConn(key)
		if conn == nil {
			continue
		}
		conn.Send(consumer.ctx, newUpdateConnCmdMsg())
	}
	return nil
}
func newUpdateConnCmdMsg() *pbmessage.PbMessage {
	return &pbmessage.PbMessage{
		MsgType: pbmessage.PbMessageType_UpdateConn,
	}
}

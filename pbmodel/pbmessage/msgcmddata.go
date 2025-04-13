/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-03 14:43:39
 */

package pbmessage

import "encoding/json"

// 二级消息使用JSON进行编码

// 更新连接信息
type UpdateConnMsg struct {
	//需要更新的客户端端id
	Keys []string
}

func NewUpdateConnPbMessage(keys []string) *PbMessage {
	data, _ := json.Marshal(newUpdateConnMsg(keys))
	return &PbMessage{
		MsgType: PbMessageType_UpdateConn,
		Data:    data,
	}
}

func newUpdateConnMsg(keys []string) *UpdateConnMsg {
	return &UpdateConnMsg{
		Keys: keys,
	}
}

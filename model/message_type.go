/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 17:25:17
 */

package model

type MessageType uint8

var (
	MessageTypeUnknown MessageType = 0 // 未知消息，占位
	MessageTypePing    MessageType = 1 // 更新心跳
	MessageTypeClose   MessageType = 2 // 关闭连接

	MessageTypePrivate MessageType = 101 // 私聊消息
	MessageTypeGroup   MessageType = 102 // 群聊消息

	ChatMessageTypeMapTypeToString = map[MessageType]string{
		MessageTypeUnknown: "message_type_unknown",
		MessageTypePing:    "message_type_ping",
		MessageTypeClose:   "message_type_close",
		MessageTypePrivate: "message_type_private",
		MessageTypeGroup:   "message_type_group",
	}
	ChatMessageTypeMapStringToType = map[string]MessageType{
		"message_type_unknown": MessageTypeUnknown,
		"message_type_ping":    MessageTypePing,
		"message_type_close":   MessageTypeClose,
		"message_type_private": MessageTypePrivate,
		"message_type_group":   MessageTypeGroup,
	}
)

func (mt *MessageType) String() string {
	typeName, ok := ChatMessageTypeMapTypeToString[*mt]
	if ok {
		return typeName
	}
	return ChatMessageTypeMapTypeToString[MessageTypeUnknown]
}

func TransferToMessageType(typeName string) MessageType {
	typeValue, ok := ChatMessageTypeMapStringToType[typeName]
	if ok {
		return typeValue
	}
	return MessageTypeUnknown
}

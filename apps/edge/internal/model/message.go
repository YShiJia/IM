/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-12 23:53:20
 */

package model

type MessageType string

const (
	MessageTypePing    MessageType = "MessageTypePing"    // 联通性测试，一般用于heartbeat
	MessageTypeClose   MessageType = "MessageTypeClose"   // 关闭连接
	MessageTypePrivate MessageType = "MessageTypePrivate" // 私聊消息
	MessageTypeGroup   MessageType = "MessageTypeGroup"   // 群聊消息
)

type ContentType string

const (
	ContentTypeText  ContentType = "ContentTypeText"  // 文本
	ContentTypeFile  ContentType = "ContentTypeFile"  // 文件
	ContentTypeImage ContentType = "ContentTypeImage" // 图片
	ContentTypeVideo ContentType = "ContentTypeVideo" // 视频
)

type Content struct {
	ContentType ContentType
	FileHash    string `json:"file_hash"` // 文件hash
	Text        string `json:"text"`      // 消息文本内容
}

type Message struct {
	MessageType `json:"message_type"` // 消息类型
	From        string                `json:"from"`      // 来自于谁
	To          string                `json:"to"`        // 发往谁
	SendTime    int64                 `json:"send_time"` // 发送方发送消息时的时间戳，用于保证消息相对有序性
	Content     Content               `json:"content"`   // 消息内容
}

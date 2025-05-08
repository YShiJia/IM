/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 18:12:19
 */

package ext

import "github.com/YShiJia/IM/model"

// Message 消息
type Message struct {
	ID          uint              `json:"id"`
	Type        model.MessageType `json:"type"`         // 消息类型
	ContentType model.ContentType `json:"content_type"` // 内容类型
	From        string            `json:"from"`         // 发送用户UID
	To          string            `json:"to"`           // 接收用户UID/群聊UID
	SendTime    int64             `json:"send_time"`    // 发送方发送消息时的时间戳，用于保证消息相对有序性
	Content     []byte            `json:"content"`      // 消息内容
}

// FileContent 文件消息内容
type FileContent struct {
	FileName string   `json:"file_name"` // 文件名
	Size     int64    `json:"size"`      // 文件大小
	Hashs    []string `json:"hashs"`     // hash列表
}

// TextContent 文本消息内容
type TextContent struct {
	Text string `json:"text"` // 文本内容
}

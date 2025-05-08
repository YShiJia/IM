/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-05 17:04:12
 */

package model

type ContentType uint8

var (
	ContentTypeUnknown ContentType = 0 // 未知内容类型
	ContentTypeText    ContentType = 1 // 文本内容
	ContentTypeImage   ContentType = 2 // 图片内容
	ContentTypeFile    ContentType = 3 // 文件内容
)

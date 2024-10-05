/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-03 14:43:39
 */

package pbmessage

// 二级消息使用JSON进行编码

// 更新连接信息
type UpdateConnMsg struct {
	//需要更新的客户端端id
	Keys []string
}

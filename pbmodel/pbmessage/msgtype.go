/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-03 13:43:26
 */

package pbmessage

type PbMessageType = uint32

// CmdMsg Type
const (
	PbMessageType_UpdateConn PbMessageType = 1001 + iota
)

// CommonMsg Type
const (
	PbMessageType_Transfer PbMessageType = 2001 + iota
)

// ClientMsg Type
const (
	PbMessageType_Ping PbMessageType = 3001 + iota
	PbMessageType_Chat PbMessageType = 3001 + iota
)

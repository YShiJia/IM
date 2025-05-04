/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-04 15:52:42
 */

package model

var MessageModels = []interface{}{
	&File{},
	&Group{},
	&User{},
	&Friend{},
	&FriendGroup{},
	&GroupMember{},
	&PrivateMessage{},
	&GroupMessage{},
}

func GetMessageModels() []interface{} {
	return MessageModels
}

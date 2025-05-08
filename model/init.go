/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-04 15:52:42
 */

package model

var MessageModels = []interface{}{
	&Group{},
	&User{},
	&GroupMember{},
	&PrivateMessage{},
	&GroupMessage{},
}

func GetMessageModels() []interface{} {
	return MessageModels
}

var FileModels = []interface{}{
	&File{},
	&FileSlice{},
}

func GetFileModels() []interface{} {
	return FileModels
}

var SocialModels = []interface{}{
	&Friend{},
	&FriendGroup{},
	&Group{},
	&User{},
	&GroupMember{},
	&PrivateMessage{},
	&GroupMessage{},
}

func GetSocialModels() []interface{} {
	return SocialModels
}

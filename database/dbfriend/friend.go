/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-16 18:01:50
 */

package dbfriend

var FriendTableName = "im_friend"

type Friend struct {
	Id        int   `json:"id" gorm:"column:id"`
	UserId    int   `json:"user_id" gorm:"column:user_id"`
	FriendId  int   `json:"friend_id" gorm:"column:friend_id"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	DeletedAt int64 `json:"deleted_at" gorm:"column:deleted_at"`
}

func (f *Friend) TableName() string {
	return "im_friend"
}

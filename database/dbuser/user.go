/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-07 22:04:51
 */

package dbuser

const (
	// 普通状态
	StatusNormal = 0
	// 冻结状态
	StatusBlock = 1
	// 删除状态
	StatusDelete = 2
)

var UserTableName = "im_user"

type User struct {
	Id        int    `json:"id" gorm:"column:id"`
	SocialId  string `json:"social_id" gorm:"column:social_id"`
	Nickname  string `json:"nickname" gorm:"column:nickname"`
	Password  string `json:"password" gorm:"column:password""`
	Email     string `json:"email" gorm:"column:email"`
	Status    int    `json:"status" gorm:"column:status"`
	Gender    int    `json:"gender" gorm:"column:gender"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at"`
	DeletedAt int64  `json:"deleted_at" gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "im_user"
}

// Check 检查字段是否完整有效
func (u *User) Check() bool {

	return true
}

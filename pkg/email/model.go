/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-14 17:39:22
 */

package email

import "time"

type EmailConfig struct {
	Username         string
	Password         string
	Host             string
	Port             string
	Expiration       time.Duration
	MaxClient        int
	TeamName         string
	CodeLen          int
	ServerExpiration time.Duration
	Interval         time.Duration
}

var EmailSuffix = []string{
	"@gmail.com",
	"@qq.com",
	"@163.com",
	"@yahoo.com",
	"@sina.com",
	"@126.com",
	"@outlook.com",
	"@yeah.net",
	"@foxmail.com",
}

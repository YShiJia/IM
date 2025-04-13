/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-14 17:42:35
 */

package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
)

func GetEmailPool(emailSetting *EmailConfig) (*email.Pool, error) {
	//验证码过期时间
	emailSetting.Expiration *= time.Second
	//邮件服务调用超时时间
	emailSetting.ServerExpiration *= time.Second
	//重复申请验证码间隔
	emailSetting.Interval *= time.Second
	//邮箱服务验证信息
	auth := smtp.PlainAuth(
		"",
		emailSetting.Username,
		emailSetting.Password,
		emailSetting.Host,
	)
	addr := fmt.Sprintf(
		"%s:%s",
		emailSetting.Host,
		emailSetting.Port,
	)
	emailPool, err := email.NewPool(addr, emailSetting.MaxClient, auth)
	if err != nil {
		return nil, err
	}
	return emailPool, nil
}

func NewEmail(emailSetting EmailConfig, html []byte, to ...string) email.Email {
	return email.Email{
		ReplyTo: []string{},
		From:    emailSetting.Username,                               //发件人邮箱
		To:      to,                                                  //收件人
		Subject: fmt.Sprintf("%s%s", emailSetting.TeamName, "平台验证码"), //主体
		HTML:    html,
	}
}

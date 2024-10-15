package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/YShiJia/IM/apps/user/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/user/rpc/user"
	"github.com/YShiJia/IM/pkg/email"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var RedisEmailVerifyCodePrefix = "redis:im:user:rpc:emailVerifyCode"

type EmailVerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailVerifyCodeLogic {
	return &EmailVerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailVerifyCodeLogic) EmailVerifyCode(in *user.EmailVerifyCodeReq) (*user.EmailVerifyCodeResp, error) {
	/*
		1. 检查email格式
		2. 查询Email是否已被绑定
		3. 生成验证码html页面 + 将验证码存到redis中
		4. 发送邮件
	*/
	if !verifyEmail(in.Email) {
		return nil, fmt.Errorf("邮箱格式不正确")
	}

	userEntity, err := l.svcCtx.UserDb.FindUserByEmail(in.Email)
	if err != nil {
		return nil, fmt.Errorf("数据库错误")
	}
	if userEntity != nil {
		return nil, fmt.Errorf("邮箱已被绑定")
	}
	//生成验证码
	code := genRandomCode(6)

	//将验证码存到redis中
	if err = l.svcCtx.Redis.SetCtx(
		l.ctx,
		fmt.Sprintf("%s:%s", RedisEmailVerifyCodePrefix, in.Email),
		code); err != nil {
	}

	//生成html页面
	htmlPage, err := render(email.VerifyCodeTemplate, email.EmailParam{
		ConfirmCode:   code,
		Expiration:    int(l.svcCtx.Config.EmailConfig.Expiration),
		OperationType: "注册验证",
		TeamName:      l.svcCtx.Config.EmailConfig.TeamName,
	})
	if err != nil {
		return nil, fmt.Errorf("html页面渲染失败 err：%w", err)
	}

	//向目标邮箱发送验证码邮件
	e := email.NewEmail(l.svcCtx.Config.EmailConfig, []byte(htmlPage), in.Email)
	if err = l.svcCtx.EmailPool.Send(&e, l.svcCtx.Config.EmailConfig.ServerExpiration*time.Second); err != nil {
		return nil, fmt.Errorf("邮件发送失败 err：%w", err)
	}
	return &user.EmailVerifyCodeResp{
		Result: "邮件发送成功",
	}, nil
}

func verifyEmail(targetEmail string) bool {
	//数据正常性验证
	//格式验证
	var flag bool
	for _, suffix := range email.EmailSuffix {
		ok := strings.HasSuffix(targetEmail, suffix)
		if ok {
			flag = true
		}
	}
	return flag
}

// 生成随机验证码
func genRandomCode(codeLen int) string {
	builder := strings.Builder{}
	for i := 0; i < codeLen; i++ {
		builder.WriteString(strconv.Itoa(rand.Int() % 10))
	}
	return builder.String()
}

// 数据渲染
func render(frameworks string, param any) (string, error) {
	tpl := template.New("")
	tpl, err := tpl.Parse(frameworks)
	if err != nil {
		return "", err
	}
	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, param)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

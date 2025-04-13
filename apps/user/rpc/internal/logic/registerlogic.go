package logic

import (
	"context"
	"fmt"
	"github.com/YShiJia/IM/apps/user/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/user/rpc/user"
	"github.com/YShiJia/IM/database/dbuser"
	"github.com/YShiJia/IM/pkg/jwt"
	"github.com/jinzhu/copier"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	/*
		1. 从 redis取出对应验证码
		2. 判断验证码是否有效
		3. 创建用户账号
		4. 返回token
	*/
	code, err := l.svcCtx.Redis.GetCtx(l.ctx, fmt.Sprintf("%s%s", RedisEmailVerifyCodePrefix, in.User.Email))
	if err != nil {
		return nil, fmt.Errorf("验证码校验失败 %w", err)
	}
	if code != in.VerifyCode {
		return nil, fmt.Errorf("验证码错误")
	}
	var DBUser dbuser.User
	copier.Copy(&DBUser, in.User)

	//TODO 需要生成唯一的社交ID

	// TODO 这里可以使用MD5加密
	DBUser.Password = in.Password
	if err = l.svcCtx.UserDb.CreateUser(&DBUser); err != nil {
		return nil, fmt.Errorf("创建用户失败")
	}
	token, err := jwt.GetJwtToken(
		l.svcCtx.Config.JwtAuth.AccessSecret,
		time.Now().Unix(),
		int64(time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire)*time.Second),
		DBUser.SocialId,
	)
	if err != nil {
		return nil, fmt.Errorf("生成token失败")
	}

	return &user.RegisterResp{
		Token:  token,
		Expire: l.svcCtx.Config.JwtAuth.AccessExpire,
		User:   in.User,
	}, nil
}

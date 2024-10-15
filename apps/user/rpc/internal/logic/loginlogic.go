package logic

import (
	"context"
	"fmt"
	"github.com/YShiJia/IM/pkg/jwt"
	"github.com/jinzhu/copier"
	"time"

	"github.com/YShiJia/IM/apps/user/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	/*
		1. 根据socialId查询用户数据
		2. 校验密码是正确
		3. 生成token返回
	*/
	DBuser, err := l.svcCtx.UserDb.GetUserBySocialId(in.SocialId)
	if err != nil {
		return nil, err
	}
	if DBuser.Password != in.Password {
		return nil, fmt.Errorf("密码错误")
	}
	token, err := jwt.GetJwtToken(
		l.svcCtx.Config.JwtAuth.AccessSecret,
		time.Now().Unix(),
		int64(time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire)*time.Second),
		in.SocialId,
	)
	if err != nil {
		return nil, fmt.Errorf("生成token失败")
	}
	userEntity := &user.UserEntity{}
	copier.Copy(userEntity, DBuser)
	return &user.LoginResp{
		Token:  token,
		Expire: l.svcCtx.Config.JwtAuth.AccessExpire,
		User:   userEntity,
	}, nil
}

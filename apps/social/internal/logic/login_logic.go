package logic

import (
	"context"
	"errors"
	conf "github.com/YShiJia/IM/apps/social/internal/config"
	"github.com/YShiJia/IM/apps/social/internal/dao/db"
	"gorm.io/gorm"
	"time"

	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"

	"github.com/YShiJia/IM/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userM, err := db.User.GetByUID(req.UID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &types.LoginResp{Message: "用户不存在"}, err
		}
		return &types.LoginResp{Message: "系统内部错误，登录失败"}, err
	}
	if userM.Password != EncryptPasswordByMD5(req.Password) {
		return &types.LoginResp{Message: "密码错误"}, err
	}
	//token 生成
	token, err := jwt.GetJwtToken(conf.Conf.AuthConf.AccessSecret, time.Now().Unix(), conf.Conf.AuthConf.AccessExpire, userM.UID)
	if err != nil {
		return &types.LoginResp{Message: "系统内部错误，登录失败"}, err
	}

	return &types.LoginResp{
		User: types.User{
			UID:      userM.UID,
			Avatar:   userM.Avatar,
			Username: userM.Name,
			Email:    userM.Email,
			Age:      userM.Age,
			Sign:     userM.Sign,
			Gender:   userM.Gender,
		},
		Token:   token,
		Message: "登录成功",
	}, nil
}

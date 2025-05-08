package logic

import (
	"context"
	"github.com/YShiJia/IM/apps/social/internal/dao/db"
	"github.com/YShiJia/IM/lib/nanoid"
	"github.com/YShiJia/IM/model"

	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	userM := &model.User{
		UID:      nanoid.User.Stand(),
		Password: EncryptPasswordByMD5(req.Password),
		Name:     req.Username,
		Email:    req.Email,
		Source:   model.SourceUserRegister,
	}
	userM, err = db.User.Create(userM)
	if err != nil {
		return &types.RegisterResp{Message: "系统内部错误，注册失败"}, err
	}
	return &types.RegisterResp{
		User: types.User{
			UID:      userM.UID,
			Avatar:   userM.Avatar,
			Username: userM.Name,
			Email:    userM.Email,
			Age:      userM.Age,
			Sign:     userM.Sign,
			Gender:   userM.Gender,
		},
		Password: req.Password,
		Message:  "注册成功",
	}, nil
}

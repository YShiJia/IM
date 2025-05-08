package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeUserInfoLogic {
	return &ChangeUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeUserInfoLogic) ChangeUserInfo(req *types.ChangeUserInfoReq) (resp *types.ChangeUserInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

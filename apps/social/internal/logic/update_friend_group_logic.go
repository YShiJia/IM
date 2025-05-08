package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFriendGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendGroupLogic {
	return &UpdateFriendGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFriendGroupLogic) UpdateFriendGroup(req *types.UpdateFriendGroupReq) (resp *types.UpdateFriendGroupResp, err error) {
	// todo: add your logic here and delete this line

	return
}

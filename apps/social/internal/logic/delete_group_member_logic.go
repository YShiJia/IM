package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupMemberLogic {
	return &DeleteGroupMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupMemberLogic) DeleteGroupMember(req *types.DeleteGroupMemberReq) (resp *types.DeleteGroupMemberResp, err error) {
	// todo: add your logic here and delete this line

	return
}

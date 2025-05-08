package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DismissGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDismissGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DismissGroupLogic {
	return &DismissGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DismissGroupLogic) DismissGroup(req *types.DismissGroupReq) (resp *types.DismissGroupResp, err error) {
	// todo: add your logic here and delete this line

	return
}

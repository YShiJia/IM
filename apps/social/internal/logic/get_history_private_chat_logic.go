package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryPrivateChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHistoryPrivateChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryPrivateChatLogic {
	return &GetHistoryPrivateChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistoryPrivateChatLogic) GetHistoryPrivateChat(req *types.GetHistoryPrivateChatReq) (resp *types.GetHistoryPrivateChatResp, err error) {
	// todo: add your logic here and delete this line

	return
}

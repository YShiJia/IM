package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryGroupChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHistoryGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryGroupChatLogic {
	return &GetHistoryGroupChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistoryGroupChatLogic) GetHistoryGroupChat(req *types.GetHistoryGroupChatReq) (resp *types.GetHistoryGroupChatResp, err error) {
	// todo: add your logic here and delete this line

	return
}

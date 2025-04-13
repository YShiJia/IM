package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/status/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/status/rpc/statusclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientMsgSyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientMsgSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientMsgSyncLogic {
	return &ClientMsgSyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientMsgSyncLogic) ClientMsgSync(in *statusclient.ClientMsgSyncRequest) (*statusclient.ClientMsgSyncResponse, error) {
	// todo: add your logic here and delete this line

	return &statusclient.ClientMsgSyncResponse{}, nil
}

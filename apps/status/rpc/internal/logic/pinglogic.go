package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/status/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/status/rpc/statusclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *statusclient.Request) (*statusclient.Response, error) {
	// todo: add your logic here and delete this line

	return &statusclient.Response{}, nil
}

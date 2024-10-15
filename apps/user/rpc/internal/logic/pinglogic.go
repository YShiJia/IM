package logic

import (
	"context"
	"fmt"

	"github.com/YShiJia/IM/apps/user/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/user/rpc/user"

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

func (l *PingLogic) Ping(in *user.Request) (*user.Response, error) {
	return &user.Response{
		Pong: fmt.Sprintf(l.svcCtx.Config.Name+" Response: ", in.Ping),
	}, nil
}

package user

import (
	"context"
	"github.com/YShiJia/IM/apps/status/rpc/statusclient"

	"github.com/YShiJia/IM/apps/status/api/internal/svc"
	"github.com/YShiJia/IM/apps/status/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WsconninfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取连接信息
func NewWsconninfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsconninfoLogic {
	return &WsconninfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsconninfoLogic) Wsconninfo(req *types.WsConnInfoReq) (resp *types.WsConnResp, err error) {
	address, err := l.svcCtx.Status.ClientConnAddress(l.ctx, &statusclient.ClientConnAddressRequest{
		SocialId: req.SocialId,
	})
	if err != nil {
		return nil, err
	}
	return &types.WsConnResp{
		Address: address.Address,
	}, nil
}

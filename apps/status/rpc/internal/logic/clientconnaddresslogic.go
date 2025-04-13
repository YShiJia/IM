package logic

import (
	"context"
	"fmt"
	"github.com/YShiJia/IM/apps/status/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/status/rpc/statusclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientConnAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientConnAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientConnAddressLogic {
	return &ClientConnAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientConnAddressLogic) ClientConnAddress(in *statusclient.ClientConnAddressRequest) (*statusclient.ClientConnAddressResponse, error) {
	nodeName, err := l.svcCtx.CsHash.GetNode(l.ctx, in.SocialId)
	if err != nil {
		return nil, err
	}
	edgeServerInfo, ok := l.svcCtx.EdgeService.Get(nodeName)
	if !ok {
		return nil, fmt.Errorf("edge server not found")
	}
	return &statusclient.ClientConnAddressResponse{
		Address: edgeServerInfo.Address,
	}, nil
}

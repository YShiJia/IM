package status

import (
	"context"
	"github.com/YShiJia/IM/apps/status/rpc/statusmodel"

	"github.com/YShiJia/IM/apps/status/api/internal/svc"
	"github.com/YShiJia/IM/apps/status/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseronlineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户在线信息
func NewUseronlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseronlineLogic {
	return &UseronlineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseronlineLogic) Useronline(req *types.UserOnlineReq) (resp *types.UserOnlineResp, err error) {
	online, err := l.svcCtx.Status.UserOnline(l.ctx, &statusmodel.UserOnlineRequest{SocialId: req.SocialIds})
	if err != nil {
		l.Logger.Error("[status api] get user online error :", err.Error())
	}
	return &types.UserOnlineResp{
		SocialIds: online.SocialId,
	}, nil
}

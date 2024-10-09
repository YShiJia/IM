package logic

import (
	"context"
	"github.com/YShiJia/IM/apps/status/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/status/rpc/statusclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOnlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOnlineLogic {
	return &UserOnlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOnlineLogic) UserOnline(in *statusclient.UserOnlineRequest) (*statusclient.UserOnlineResponse, error) {
	/**
	1. 根据socialIds先获取所有用户的ids
	2. 根据ids在bitmap中相应的位置是否为1
	3. 根据在线的ids查询的socialIds
	4. 返回在线的socialIds
	*/
	ids, err := l.svcCtx.UserDb.GetIdBySocialId(in.SocialId)
	if err != nil {
		return nil, err
	}
	array, err := l.svcCtx.UserOnlineBitMap.CountArray(l.ctx, ids)
	if err != nil {
		return nil, err
	}
	socialIds, err := l.svcCtx.UserDb.GetSocialIdById(array)
	if err != nil {
		return nil, err
	}

	return &statusclient.UserOnlineResponse{
		SocialId: socialIds,
	}, nil
}

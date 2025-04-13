package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFriendLogic) DeleteFriend(in *social.DeleteFriendReq) (*social.DeleteFriendResp, error) {
	if err := l.svcCtx.FriendDB.DeleteFriend(in.FromSocialId, in.ToSocialId); err != nil {
		return nil, err
	}
	return &social.DeleteFriendResp{
		Result: "delete friend relationship success",
	}, nil
}

package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeFriendLogic {
	return &MakeFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeFriendLogic) MakeFriend(in *social.MakeFriendReq) (*social.MakeFriendResp, error) {
	if err := l.svcCtx.FriendDB.MakeFriend(in.FromSocialId, in.ToSocialId); err != nil {
		return nil, err
	}
	return &social.MakeFriendResp{
		Result: "make friend success",
	}, nil
}

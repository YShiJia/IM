package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupLogic) JoinGroup(in *social.JoinGroupReq) (*social.JoinGroupResp, error) {
	err := l.svcCtx.GroupDB.JoinGroup(in.UserSocialId, in.GroupSocialId)
	if err != nil {
		return nil, err
	}
	return &social.JoinGroupResp{
		Result: "join group success",
	}, nil
}

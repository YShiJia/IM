package logic

import (
	"context"

	"github.com/YShiJia/IM/apps/social/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuitGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuitGroupLogic {
	return &QuitGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuitGroupLogic) QuitGroup(in *social.QuitGroupReq) (*social.QuitGroupResp, error) {
	err := l.svcCtx.GroupDB.QuitGroup(in.UserSocialId, in.GroupSocialId)
	if err != nil {
		return nil, err
	}
	return &social.QuitGroupResp{
		Result: "quit group success",
	}, nil
}

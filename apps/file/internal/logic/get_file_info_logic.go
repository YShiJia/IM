package logic

import (
	"context"
	"github.com/YShiJia/IM/apps/file/internal/dao/db"
	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileInfoLogic) GetFileInfo(req *types.GetFileInfoReq) (resp *types.GetFileInfoResp, err error) {
	fileM, err := db.File.GetByHash(req.Hash)
	if err != nil {
		return nil, err
	}
	fileSlices, err := db.FileSlice.ListByFileID(fileM.ID, db.FileSlice.OrderByOrder())
	if err != nil {
		return nil, err
	}
	resp = &types.GetFileInfoResp{
		Hash:       req.Hash,
		FileSlices: []types.FileSlice{},
	}
	for _, fileSlice := range fileSlices {
		resp.FileSlices = append(resp.FileSlices, types.FileSlice{
			Order: fileSlice.Order,
			Hash:  fileSlice.Hash,
		})
	}
	return resp, nil
}

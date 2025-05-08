package logic

import (
	"context"
	"github.com/YShiJia/IM/apps/file/internal/dao/db"
	"github.com/YShiJia/IM/apps/file/internal/dao/fs"
	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"
	"github.com/YShiJia/IM/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileSliceUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileSliceUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileSliceUploadPrepareLogic {
	return &FileSliceUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileSliceUploadPrepareLogic) FileSliceUploadPrepare(req *types.FileSliceUploadPrepareReq) (resp *types.FileSliceUploadPrepareResp, err error) {
	// 创建fileM
	fileM := &model.File{
		Bucket: fs.GetBucketName(req.FileName),
		Hash:   req.Hash,
		Size:   req.Size,
		Name:   req.FileName,
	}
	fileM, err = db.File.Create(fileM)
	if err != nil {
		return nil, err
	}

	return &types.FileSliceUploadPrepareResp{
		FileName: fileM.Name,
		Size:     fileM.Size,
		Hash:     fileM.Hash,
	}, nil
}

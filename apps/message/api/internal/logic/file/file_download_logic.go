package file

import (
	"context"
	"fmt"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/db"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/fs"
	"github.com/YShiJia/IM/apps/message/api/internal/svc"
	"github.com/YShiJia/IM/apps/message/api/internal/types"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDownloadLogic {
	return &FileDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDownloadLogic) FileDownload(req *types.FileDownloadReq) (*types.FileDownloadResp, error) {
	// 先查询是否存在
	fileM, err := db.File.GetByHash(req.Hash)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Errorf("file[%s] 不存在", req.Hash)
			return nil, fmt.Errorf("file[%s] 不存在", req.Hash)
		}
		log.Errorf("根据hash[%s]获取file model 失败", req.Hash)
		return nil, fmt.Errorf("根据hash[%s]获取file model 失败", req.Hash)
	}
	data, err := fs.Minio.GetFile(l.ctx, fileM.Bucket, fileM.Hash)
	if err != nil {
		log.Errorf("下载文件[%s]失败", fileM.Hash)
		return nil, fmt.Errorf("下载文件[%s]失败", fileM.Hash)
	}

	return &types.FileDownloadResp{Data: data}, nil
}

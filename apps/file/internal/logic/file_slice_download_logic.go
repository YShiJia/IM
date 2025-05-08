package logic

import (
	"bytes"
	"context"
	"errors"
	"github.com/YShiJia/IM/apps/file/internal/dao/db"
	"github.com/YShiJia/IM/apps/file/internal/dao/fs"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"

	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileSliceDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileSliceDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileSliceDownloadLogic {
	return &FileSliceDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileSliceDownloadLogic) FileSliceDownload(req *types.FileSliceDownloadReq, w http.ResponseWriter) (resp *types.FileSliceDownloadResp, err error) {
	fileSlice, err := db.FileSlice.GetByHash(req.Hash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("file slice[hash:%s] not found", req.Hash)
		}
		log.Errorf("get file slice[hash:%s] error:%v", req.Hash, err)
		return
	}
	// 获取文件分片
	fileData, err := fs.Minio.GetFile(l.ctx, fileSlice.Bucket, fileSlice.Hash)
	if err != nil {
		log.Errorf("get file slice[hash:%s] error:%v", req.Hash, err)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)

	// 写入数据
	if _, err = io.Copy(w, bytes.NewReader(fileData)); err != nil {
		log.Errorf("Write file data error: %v", err)
		return nil, err
	}

	return nil, nil
}

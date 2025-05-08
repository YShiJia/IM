package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/YShiJia/IM/apps/file/internal/dao/db"
	"github.com/YShiJia/IM/apps/file/internal/dao/fs"
	log "github.com/sirupsen/logrus"
	"io"
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"

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

func (l *FileDownloadLogic) FileDownload(req *types.FileDownloadReq, w http.ResponseWriter) (resp *types.FileDownloadResp, err error) {

	ext := path.Ext(req.FileName)
	if ext == "" {
		log.Errorf("Missing file extension fileName:[%s]", req.FileName)
		return
	}
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	hash := strings.TrimSuffix(req.FileName, ext)

	// 查询是否有该文件
	fileM, err := db.File.GetByHash(hash)
	if err != nil {
		log.Errorf("Get file by hash error: %v", err)
		return nil, err
	}
	fileData, err := fs.Minio.GetFile(l.ctx, fileM.Bucket, fileM.Hash)
	if err != nil {
		log.Errorf("Get file by hash error: %v", err)
		return nil, err
	}

	// 设置响应头
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", req.FileName))
	w.WriteHeader(http.StatusOK)

	// 写入数据
	if _, err = io.Copy(w, bytes.NewReader(fileData)); err != nil {
		log.Errorf("Write file data error: %v", err)
		return nil, err
	}

	return &types.FileDownloadResp{}, nil
}

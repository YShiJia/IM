package logic

import (
	"bytes"
	"context"
	"errors"
	"github.com/YShiJia/IM/apps/file/internal/dao/db"
	"github.com/YShiJia/IM/apps/file/internal/dao/fs"
	"github.com/YShiJia/IM/lib/sign"
	"github.com/YShiJia/IM/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileSliceUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileSliceUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileSliceUploadLogic {
	return &FileSliceUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileSliceUploadLogic) FileSliceUpload(req *types.FileSliceUploadReq, fileData []byte) (resp *types.FileSliceUploadResp, err error) {
	// 先查询是否有这个主文件
	fileM, err := db.File.GetByHash(req.FileHash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("文件[hash:%s]不存在", req.FileHash)
		}
		log.Errorf("根据文件hash[%s]查询文件失败，err: %v", req.FileHash, err)
		return nil, err
	}
	md5, err := sign.FileSignatureByMD5(bytes.NewReader(fileData))
	if err != nil {
		log.Errorf("获取文件MD5失败，err：%v", err)
		return nil, err
	}
	fileSliceM := &model.FileSlice{
		Bucket: fileM.Bucket,
		Hash:   md5,
		Size:   req.Size,
		FileID: fileM.ID,
		Order:  req.Order,
	}
	if err = db.IMDB.Transaction(func(tx *gorm.DB) error {
		fileSliceDAO := db.NewFileSliceDAO(tx)
		fileSliceM, err = fileSliceDAO.FirstOrCreate(fileSliceM)
		if err != nil {
			log.Errorf("根据文件hash[%s]查询或创建文件失败，err: %v", req.FileHash, err)
			return err
		}
		if err = fs.Minio.UploadFile(l.ctx, fileSliceM.Bucket, fileSliceM.Hash, fileSliceM.Size, bytes.NewReader(fileData)); err != nil {
			log.Errorf("上传文件失败，err: %v", err)
			return err
		}
		return nil
	}); err != nil {
		log.Errorf("上传文件失败，err: %v", err)
		return nil, err
	}

	return &types.FileSliceUploadResp{
		Hash:  md5,
		Order: fileSliceM.Order,
	}, nil
}

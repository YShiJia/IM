package file

import (
	"bytes"
	"context"
	"fmt"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/db"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/fs"
	"github.com/YShiJia/IM/apps/message/api/internal/svc"
	"github.com/YShiJia/IM/apps/message/api/internal/types"
	"github.com/YShiJia/IM/apps/message/api/model"
	"github.com/YShiJia/IM/lib/sign"
	imModel "github.com/YShiJia/IM/model"
	log "github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq, data []byte, fileHeader *multipart.FileHeader) (resp *types.FileUploadResp, err error) {
	//判断是否已存在
	fileM, err := db.File.GetByHash(req.Hash)
	if err == nil {
		return &types.FileUploadResp{Hash: fileM.Hash}, err
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Infof("根据hash获取file model err：%v", err)
		return nil, fmt.Errorf("根据hash获取file model err：%v", err)
	}
	// 文件不存在，创建文件
	hash, err := sign.FileSignatureByMD5(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("获取文件hash值失败 err：%v", err)
	}
	// 上传文件，再创建记录
	bucketName := fs.GetBucketName(req.FileName)
	if err = fs.Minio.UploadFile(l.ctx, bucketName, hash, int64(len(data)), bytes.NewReader(data)); err != nil {
		log.Infof("文件[%s]上传失败%v", hash, err)
	}
	// 创建记录
	fileM = &model.File{
		Audit: imModel.Audit{
			Creator:   "todo",
			CreatedAt: time.Now(),
			Updater:   "todo",
			UpdatedAt: time.Now(),
		},
		Hash:   hash,
		Bucket: bucketName,
		Name:   req.FileName,
	}
	fileM, err = db.File.Create(fileM)
	if err != nil {
		log.Infof("创建文件[%s]记录失败%v", hash, err)
		return nil, fmt.Errorf("创建文件[%s]记录失败%v", hash, err)
	}
	return &types.FileUploadResp{Hash: fileM.Hash}, err
}

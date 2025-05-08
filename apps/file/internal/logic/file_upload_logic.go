package logic

import (
	"bytes"
	"context"
	"github.com/YShiJia/IM/apps/file/internal/dao/db"
	"github.com/YShiJia/IM/apps/file/internal/dao/fs"
	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"
	"github.com/YShiJia/IM/lib/sign"
	"github.com/YShiJia/IM/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq, fileData []byte) (resp *types.FileUploadResp, err error) {

	md5, err := sign.FileSignatureByMD5(bytes.NewReader(fileData))
	if err != nil {
		log.Errorf("获取文件MD5失败，err：%v", err)
		return nil, err
	}
	fileM := &model.File{
		Bucket: fs.GetBucketName(req.FileName),
		Hash:   md5,
		Size:   req.Size,
		Name:   req.FileName,
	}
	fileSliceM := &model.FileSlice{
		Bucket: fileM.Bucket,
		Hash:   fileM.Hash,
		Size:   fileM.Size,
		Order:  1,
	}
	if err = db.IMDB.Transaction(func(tx *gorm.DB) error {
		fileDAO := db.NewFileDAO(tx)
		fileSliceDAO := db.NewFileSliceDAO(tx)

		// 创建file
		fileM, err := fileDAO.Create(fileM)
		if err != nil {
			log.Errorf("创建fileM失败，err：%v", err)
			return err
		}

		fileSliceM.FileID = fileM.ID
		// 创建file_slice
		if _, err = fileSliceDAO.Create([]*model.FileSlice{fileSliceM}); err != nil {
			log.Errorf("创建file_slice失败，err：%v", err)
			return err
		}
		// 将文件数据存储到minio中
		if err = fs.Minio.UploadFile(context.TODO(), fileM.Bucket, fileM.Hash, fileM.Size, bytes.NewReader(fileData)); err != nil {
			log.Errorf("将文件数据存储到minio中失败，err：%v", err)
			return err
		}
		return nil
	}); err != nil {
		log.Errorf("存储文件事务执行失败，err：%v", err)
		return nil, err
	}

	return &types.FileUploadResp{
		Hash: md5,
		FileSlices: []types.FileSlice{
			{
				Hash:  md5,
				Order: 1,
			},
		},
	}, nil
}

package file

import (
	"context"
	"fmt"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/YShiJia/IM/apps/message/api/internal/svc"
	"github.com/YShiJia/IM/apps/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileCheckLogic {
	return &FileCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileCheckLogic) FileCheck(req *types.FileCheckReq) (resp *types.FileCheckResp, err error) {
	_, err = db.File.GetByHash(req.Hash)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &types.FileCheckResp{Exists: false}, nil
		}
		log.Errorf("根据hash[%s]查询文件数据错误, err %v", req.Hash, err)
		return nil, fmt.Errorf("根据hash[%s]查询文件数据错误, err %v", req.Hash, err)
	}

	return &types.FileCheckResp{Exists: true}, nil
}

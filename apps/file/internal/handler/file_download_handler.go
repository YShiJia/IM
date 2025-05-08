package handler

import (
	"errors"
	"gorm.io/gorm"
	"net/http"

	"github.com/YShiJia/IM/apps/file/internal/logic"
	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileDownloadLogic(r.Context(), svcCtx)
		_, err := l.FileDownload(&req, w)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				httpx.WriteJson(w, http.StatusNotFound, nil)
			}
			httpx.ErrorCtx(r.Context(), w, err)
		}
	}
}

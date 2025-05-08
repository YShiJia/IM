package handler

import (
	"net/http"

	"github.com/YShiJia/IM/apps/file/internal/logic"
	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileSliceDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileSliceDownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileSliceDownloadLogic(r.Context(), svcCtx)
		_, err := l.FileSliceDownload(&req, w)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
	}
}

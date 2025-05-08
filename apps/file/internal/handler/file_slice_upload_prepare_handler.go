package handler

import (
	"net/http"

	"github.com/YShiJia/IM/apps/file/internal/logic"
	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileSliceUploadPrepareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileSliceUploadPrepareReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileSliceUploadPrepareLogic(r.Context(), svcCtx)
		resp, err := l.FileSliceUploadPrepare(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

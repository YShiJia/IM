package file

import (
	"net/http"

	"github.com/YShiJia/IM/apps/message/api/internal/logic/file"
	"github.com/YShiJia/IM/apps/message/api/internal/svc"
	"github.com/YShiJia/IM/apps/message/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileCheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewFileCheckLogic(r.Context(), svcCtx)
		resp, err := l.FileCheck(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

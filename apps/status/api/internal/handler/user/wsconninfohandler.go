package user

import (
	"net/http"

	"github.com/YShiJia/IM/apps/status/api/internal/logic/user"
	"github.com/YShiJia/IM/apps/status/api/internal/svc"
	"github.com/YShiJia/IM/apps/status/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取连接信息
func WsconninfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsConnInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewWsconninfoLogic(r.Context(), svcCtx)
		resp, err := l.Wsconninfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package status

import (
	"net/http"

	"github.com/YShiJia/IM/apps/status/api/internal/logic/status"
	"github.com/YShiJia/IM/apps/status/api/internal/svc"
	"github.com/YShiJia/IM/apps/status/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户在线信息
func UseronlineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOnlineReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := status.NewUseronlineLogic(r.Context(), svcCtx)
		resp, err := l.Useronline(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

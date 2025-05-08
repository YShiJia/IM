package handler

import (
	"net/http"

	"github.com/YShiJia/IM/apps/social/internal/logic"
	"github.com/YShiJia/IM/apps/social/internal/svc"
	"github.com/YShiJia/IM/apps/social/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteGroupMemberHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteGroupMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteGroupMemberLogic(r.Context(), svcCtx)
		resp, err := l.DeleteGroupMember(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

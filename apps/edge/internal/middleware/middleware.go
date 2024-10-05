/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 20:06:50
 */

package middleware

import (
	"github.com/YShiJia/IM/apps/edge/internal/svc"
	"github.com/YShiJia/IM/lib/webscoket"
	"github.com/YShiJia/IM/pkg/jwt"
	"net/http"
)

func WsConnJwtAuthorize(svcCtx *svc.ServiceContext) webscoket.HandleFunc {
	return func(ctx *webscoket.WebContext) {
		token := ctx.R.Header.Get("Authorization")
		if token == "" {
			ctx.W.WriteHeader(http.StatusUnauthorized)
			ctx.W.Write([]byte("auth information is not exist"))
			ctx.Abort()
			return
		}
		id, err := jwt.ParseJwtToken(svcCtx.Config.Auth.AccessSecret, token)
		if err != nil {
			ctx.W.WriteHeader(http.StatusUnauthorized)
			ctx.W.Write([]byte("auth information is not exist"))
			ctx.Abort()
			return
		}
		ctx.Set(webscoket.WsConnectionIdentity, id)
	}
}

/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 20:06:50
 */

package middleware

import (
	"fmt"
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	libWs "github.com/YShiJia/IM/lib/websocket"
	"github.com/YShiJia/IM/pkg/jwt"
	"github.com/gorilla/websocket"
)

// JwtAuthorize 认证并存储用户身份信息
func JwtAuthorize(ctx *libWs.WebContext) {
	token := ctx.R.Header.Get(conf.Authorization)
	if token == "" {
		ctx.Conn().Send(websocket.TextMessage, []byte(fmt.Sprintf("header %s is not exist", conf.Authorization)))
		ctx.Abort()
		return
	}
	// uid 为用户uid
	uid, err := jwt.ParseJwtToken(conf.Conf.AuthConf.AccessSecret, token)
	if err != nil {
		ctx.Conn().Send(websocket.TextMessage, []byte("auth information is invalid"))
		ctx.Abort()
		return
	}
	ctx.Set(conf.AuthIMUserUID, uid)
}

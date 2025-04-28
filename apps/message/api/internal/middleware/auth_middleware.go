package middleware

import (
	"github.com/YShiJia/IM/apps/define"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/pkg/jwt"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		writeErr := func(w http.ResponseWriter) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}

		token := r.Header.Get(define.Authorization)
		if token == "" {
			writeErr(w)
			return
		}
		// 无状态服务，无需存储id
		_, err := jwt.ParseJwtToken(conf.Conf.AuthConf.AccessSecret, token)
		if err != nil {
			writeErr(w)
			return
		}
		next(w, r)
	}
}

/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-06 15:29:36
 */

package middleware

import "net/http"

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // 允许所有域访问
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // 允许的请求方法
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // 允许的请求头

		// 对于 OPTIONS 请求，直接返回
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

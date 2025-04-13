/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 20:17:30
 */

package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

const identifier = "uid"

func GetJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	//下面两个字段是内置的，会自动校验过期时间
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[identifier] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseJwtToken(secretKey string, payload string) (string, error) {
	// 解析 JWT
	token, err := jwt.Parse(payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	// 验证并获取声明
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims[identifier].(string), nil
	}

	return "", fmt.Errorf("invalid JWT token")
}

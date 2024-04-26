package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var ScreenKey string

type Payload struct {
	Id      uint   `json:"id"`
	Account string `json:"account"`
	Role    uint   `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(claims Payload) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
		IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
		NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(ScreenKey))
	return s, err
}

func ParseJwt(tokenstring string) (*Payload, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ScreenKey), nil
	})

	if claims, ok := t.Claims.(*Payload); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"mini-tiktok-hanyongyan/pkg/config"
	"time"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

// CreateToken 生成 token
func CreateToken(userId int64) (string, error) {
	expiretime := time.Now().Add(24 * 7 * time.Hour)
	nowTime := time.Now()

	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiretime.Unix(),
			Issuer:    "mini-tiktok",
			IssuedAt:  nowTime.Unix(),
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString([]byte(config.SecretKey))
}

// CheckToken 解析 token
func CheckToken(token string) (*Claims, bool) {
	tokenObj, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	if key, _ := tokenObj.Claims.(*Claims); tokenObj.Valid {
		return key, true
	}
	return nil, false
}

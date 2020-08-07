package util

import (
	"go_web/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.Set.JwtKey)

type Claims struct {
	UserID   uint   `json:"id"`
	Username string `json:"username"`
	RoleID   uint64 `json:"role_id"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, role_id uint64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		id,
		username,
		role_id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-web",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

package util

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

var (
	ExpireTime = 30 * 24 * time.Hour
)


type Claims struct {
	UserId uint `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}


func GenerateToken(userid uint, username string) string  {
	nowTime := time.Now()
	expireTime := nowTime.Add(ExpireTime)

	user := Claims{
		userid,
		username,
		jwt.StandardClaims {
			ExpiresAt: expireTime.Unix(),
			Issuer: "giligili",
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	if token, err := claims.SignedString(secretKey); err != nil {
		panic(err)
		return ""
	} else {
		return token
	}
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

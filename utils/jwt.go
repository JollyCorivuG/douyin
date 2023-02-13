package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// token持续时间
const durationTime = time.Hour * 7 * 24

// 加密key
var jwtSecretKey = []byte("jollycorivug.jhc")

type Cliams struct {
	UserId int64
	jwt.StandardClaims
}

// 颁发token
func ReleaseToken(userId int64) (string, error) {
	expirationTime := time.Now().Add(durationTime)
	cliams := &Cliams{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "jhc",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*Cliams, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Cliams{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	cliams, ok := token.Claims.(*Cliams)
	if ok && token.Valid {
		return cliams, nil
	}
	return nil, errors.New("invalid token")
}

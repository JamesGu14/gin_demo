package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var app_key = []byte("my-secrete")

type UserInfo struct {
	StudentId uint
}

type Claims struct {
	UserInfo
	jwt.StandardClaims
}

// Issue token
func GenerateJwtToken(studentId uint) string {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin_demo",
			Subject:   "Student Token",
		},
	}
	claims.StudentId = studentId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(app_key)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

// Parse token
func ParseToken(tokenString string) (*UserInfo, bool) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return app_key, nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}
	return &Claims.UserInfo, true
}

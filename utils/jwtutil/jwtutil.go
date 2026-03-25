package jwtutil

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"krm-backend/config"
	"krm-backend/utils/logs"
	"time"
)

var jwtSignKey = []byte(config.JwtSignKey)

// 1,自定义声明类型
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 2.封装生产token的函数
func GenToken(username string) (string, error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.JwtExpTime) * time.Minute)), //设置token有效期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "hhhh",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)
	return ss, err
}

// 3.解析token
func ParseToken(tokenStr string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		return jwtSignKey, nil
	})
	if err != nil {
		//解析token失败
		logs.Error(nil, "解析token失败") //就是这根本就不是一个token
		return nil, err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//token合法
		return claims, nil
	} else {
		//token不合法，是一个token但是有问题（可能是过期了等等）
		logs.Warning(nil, "token不合法")
		return nil, errors.New("token不合法")
	}
}

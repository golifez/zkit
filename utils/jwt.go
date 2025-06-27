package utils

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type jwt struct {
}

// NewJwt 生成jwt
func NewJwt() *jwt {
	return &jwt{}
}

// 获取token
func (j *jwt) GetToken(claims map[string]interface{}, key string) (token string) {
	jclm := jwtv5.MapClaims{}

	for k, v := range claims {
		if v != nil {
			jclm[k] = v
		}
	}
	tokenv := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, jclm)
	tokenString, _ := tokenv.SignedString([]byte(key))
	fmt.Println(tokenString)
	return tokenString
}

// 解析token
func (j *jwt) ParseToken(token string, key string) (claims jwtv5.MapClaims, err error) {
	tokenv, err := jwtv5.Parse(token, func(token *jwtv5.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			return nil, errors.Unauthorized("unauthorized", "unexpected signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		//细节处理token
		if errors.Is(err, jwtv5.ErrTokenExpired) {
			return nil, errors.Unauthorized("unauthorized", "token expired")
		}
		if errors.Is(err, jwtv5.ErrTokenMalformed) {
			return nil, errors.Unauthorized("unauthorized", "token malformed")
		}
		return nil, errors.Unauthorized("无效的token", err.Error())
	}
	if claims, ok := tokenv.Claims.(jwtv5.MapClaims); ok && tokenv.Valid {
		return claims, nil
	}
	return nil, errors.Unauthorized("invalid token", "invalid token")
}

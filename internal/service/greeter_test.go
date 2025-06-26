package service

import (
	"fmt"
	"testing"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func TestGree(t *testing.T) {
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, jwtv5.MapClaims{
		"sub":    "1234567890",
		"name":   "zxb",
		"passwd": "abc123321",
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("test"))
	fmt.Println(tokenString)
}

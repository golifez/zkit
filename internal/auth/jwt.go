package auth

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"

	"github.com/duke-git/lancet/random"

	"github.com/duke-git/lancet/v2/cryptor"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// 获取token
func GetToken(uid string, key string) (token string) {
	tokenv := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, jwtv5.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, _ := tokenv.SignedString([]byte(key))
	fmt.Println(tokenString)
	return tokenString
}

// 解析token
func ParseToken(token string, key string) (claims jwtv5.MapClaims, err error) {
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

// 加密密码
func Encryption(passwd string, key string) (encrypteds string) {
	encrypted := cryptor.AesEcbEncrypt([]byte(passwd), []byte(key))
	// 将二进制数据转换为Base64编码的字符串
	return base64.StdEncoding.EncodeToString(encrypted)
}

// 解密密码
func Decrypt(passwd string, key string) (decrypteds string, err error) {
	// 将Base64编码的字符串解码为二进制数据
	encryptedBytes, err := base64.StdEncoding.DecodeString(passwd)
	if err != nil {
		fmt.Printf("base64解码失败:%s", err)
		return "", err
	}
	decrypted := cryptor.AesEcbDecrypt(encryptedBytes, []byte(key))
	fmt.Printf("解密后的密码是:%s", string(decrypted))
	return string(decrypted), nil
}

// 生成UID

func GenUid() string {
	uuid, err := random.UUIdV4()
	if err != nil {
		return ""
	}
	fmt.Println(uuid)
	return uuid
}

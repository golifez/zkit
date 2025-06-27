package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/duke-git/lancet/v2/cryptor"
)

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

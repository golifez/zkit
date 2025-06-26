package auth

import (
	"fmt"
	"testing"
)

func TestEncryption(t *testing.T) {
	originalPassword := "R8b3tEw28yWQ2cHNxQN3GQ=="
	key := "yLcTKVyWyssTsTsg"
	// 加密
	encrypted := Encryption(originalPassword, key)
	fmt.Printf("原始密码: %s\n", originalPassword)
	fmt.Printf("加密后密码: %s\n", encrypted)

	// 解密验证
	decrypted, _ := Decrypt(encrypted, key)
	fmt.Printf("解密后密码: %s\n", decrypted)

	// 验证解密结果是否正确
	if decrypted != originalPassword {
		t.Errorf("解密失败: 期望 %s, 得到 %s", originalPassword, decrypted)
	}
}

func TestDecryptDatabasePassword(t *testing.T) {
	// 测试数据库里的密码
	databasePassword := "R8b3tEw28yWQ2cHNxQN3GQ=="
	key := "yLcTKVyWyssTsTsg"

	fmt.Printf("数据库密码: %s\n", databasePassword)
	fmt.Printf("解密密钥: %s\n", key)

	decrypted, err := Decrypt(databasePassword, key)
	if err != nil {
		t.Errorf("解密失败: %v", err)
		return
	}

	fmt.Printf("解密结果: '%s' (长度: %d)\n", decrypted, len(decrypted))

	// 测试用相同密钥重新加密
	reEncrypted := Encryption(decrypted, key)
	fmt.Printf("重新加密结果: %s\n", reEncrypted)

	// 验证是否一致
	if reEncrypted != databasePassword {
		t.Errorf("重新加密结果不一致: 期望 %s, 得到 %s", databasePassword, reEncrypted)
	}
}

func TestGenUid(t *testing.T) {
	uid := GenUid()
	fmt.Println("uid是", uid)
}

// func TestEncryption(t *testing.T) {
// 	originalPassword := "abc123"
// 	key := "yLcTKVyWyssTsTsg"
// 	// 加密
// 	encrypted := Encryption(originalPassword, key)
// 	fmt.Printf("原始密码: %s\n", originalPassword)
// 	fmt.Printf("加密后密码: %s\n", encrypted)

// 	// 解密验证
// 	decrypted, _ := Decrypt(encrypted, key)
// 	fmt.Printf("解密后密码: %s\n", decrypted)

// 	// 验证解密结果是否正确
// 	if decrypted != originalPassword {
// 		t.Errorf("解密失败: 期望 %s, 得到 %s", originalPassword, decrypted)
// 	}
// }

func TestEnAndDePassword(t *testing.T) {
	// 测试数据库里的密码
	Password := "abc123"
	key := "yLcTKVyWyssTsTsg"

	fmt.Printf("原始密码: %s\n", Password)
	fmt.Printf("加密密钥: %s\n", key)

	encrypted := Encryption(Password, key)
	fmt.Printf("加密后密码: %s\n", encrypted)

	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		t.Errorf("解密失败: %v", err)
		return
	}

	fmt.Printf("解密结果: '%s' (长度: %d)\n", decrypted, len(decrypted))

	// 测试用相同密钥重新加密
	reEncrypted := Encryption(decrypted, key)
	fmt.Printf("重新加密结果: %s\n", reEncrypted)

	// 验证是否一致
	if reEncrypted != encrypted {
		t.Errorf("重新加密结果不一致: 期望 %s, 得到 %s", encrypted, reEncrypted)
	}
}

func TestVerifyDatabasePassword(t *testing.T) {
	// 数据库里的密码
	databasePassword := "R8b3tEw28yWQ2cHNxQN3GQ=="
	key := "yLcTKVyWyssTsTsg"

	fmt.Printf("数据库密码: %s\n", databasePassword)
	fmt.Printf("当前密钥: %s\n", key)

	// 尝试解密
	decrypted, err := Decrypt(databasePassword, key)
	if err != nil {
		t.Errorf("解密失败: %v", err)
		return
	}

	fmt.Printf("解密结果: '%s' (长度: %d)\n", decrypted, len(decrypted))

	// 如果解密出来是空的，说明这个密码不是用当前方法加密的
	if len(decrypted) == 0 {
		fmt.Println("❌ 数据库密码不是用当前加密方法加密的！")
		fmt.Println("需要重新设置数据库密码")
	} else {
		fmt.Println("✅ 数据库密码是用当前加密方法加密的")
	}

	// 测试用 abc123 加密，看看结果
	testPassword := "abc123"
	encrypted := Encryption(testPassword, key)
	fmt.Printf("'%s' 加密后: %s\n", testPassword, encrypted)

	// 比较
	if encrypted == databasePassword {
		fmt.Println("✅ 数据库密码就是 'abc123' 加密的结果")
	} else {
		fmt.Println("❌ 数据库密码不是 'abc123' 加密的结果")
	}
}

func TestParseToken(t *testing.T) {
	// 生成一个token
	token := GetToken("123", "admin")
	fmt.Printf("生成的token: %s\n", token)
	// 解析token
	claims, err := ParseToken(token, "admin")
	if err != nil {
		t.Errorf("ParseToken failed: %v", err)
		return
	}
	fmt.Printf("解析后的claims: %v\n", claims)
}

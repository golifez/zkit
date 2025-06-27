package utils

import (
	"fmt"

	"github.com/duke-git/lancet/random"
)

// 生成UID-v4

func GenUid() string {
	uuid, err := random.UUIdV4()
	if err != nil {
		return ""
	}
	fmt.Println(uuid)
	return uuid
}

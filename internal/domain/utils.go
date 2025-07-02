package domain

import (
	"time"

	pb "github.com/golifez/zkit/api/zkitauth/v1"
)

type Jwt struct {
	Key    string
	Claims map[string]interface{}
}

func NewGenTokenRequestFromApi(req *pb.GenTokenRequest) *Jwt {
	myclaims := make(map[string]interface{}, len(req.Claims))
	// 先拷贝一份原始 Claims
	for k, v := range req.Claims {
		myclaims[k] = v
	}
	//过期时间
	myclaims["exp"] = time.Now().Add(12 * time.Hour).Unix()
	// // 处理 exp 字段
	// if raw, ok := myclaims["exp"]; ok {
	// 	var expUnix int64
	// 	switch val := raw.(type) {
	// 	case int64:
	// 		expUnix = val
	// 	case float64:
	// 		// JSON 解出来的数字有可能是 float64
	// 		expUnix = int64(val)
	// 	case string:
	// 		// 字符串转数字
	// 		if parsed, err := strconv.ParseInt(val, 10, 64); err == nil {
	// 			expUnix = parsed
	// 		} else {
	// 			// 解析失败，默认 1 小时后过期
	// 			expUnix = time.Now().Add(time.Hour).Unix()
	// 		}
	// 	default:
	// 		// 其它类型，默认 1 小时后过期
	// 		expUnix = time.Now().Add(time.Hour).Unix()
	// 	}
	// 	myclaims["exp"] = expUnix
	// } else {
	// 	// 如果没有传 exp，直接设为 1 小时后
	// 	myclaims["exp"] = time.Now().Add(time.Hour).Unix()
	// }

	return &Jwt{
		Key:    req.Key,
		Claims: myclaims,
	}
}

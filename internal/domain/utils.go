package domain

import (
	pb "github.com/golifez/zkit/api/zkitauth/v1"
)

type Jwt struct {
	Key    string
	Claims map[string]interface{}
}

func NewGenTokenRequestFromApi(req *pb.GenTokenRequest) *Jwt {

	myclaims := map[string]interface{}{}
	for k, v := range req.Claims {
		myclaims[k] = v
	}
	return &Jwt{
		Key:    req.Key,
		Claims: myclaims,
	}
}

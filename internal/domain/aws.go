package domain

import (
	pb "github.com/golifez/zkit/api/aws/v1"
)

type AwsIam struct {
	AccountId string
	AccessKey string
	SecretKey string
}

func NewAddAwsIamFromApi(req *pb.AddAkSecretRequest) *AwsIam {
	return &AwsIam{
		AccountId: req.AconutId,
		AccessKey: req.AccessKey,
		SecretKey: req.SecretKey,
	}
}

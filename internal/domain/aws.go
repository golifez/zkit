package domain

import (
	pb "github.com/golifez/zkit/api/aws/v1"
)

type AwsIam struct {
	Uid       string
	AccountId string
	AccessKey string
	SecretKey string
}

func NewAddAwsIamFromApi(req *pb.AddAkSecretRequest) *AwsIam {
	return &AwsIam{
		Uid:       req.Uid,
		AccountId: req.AccountId,
		AccessKey: req.AccessKey,
		SecretKey: req.SecretKey,
	}
}

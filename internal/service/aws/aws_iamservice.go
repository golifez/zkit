package aws

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/golifez/zkit/api/aws/v1"
	bitaws "github.com/golifez/zkit/internal/biz/aws"
	"github.com/golifez/zkit/internal/domain"
)

type IamService struct {
	pb.UnimplementedIamServiceServer
	log *log.Helper
	uc  *bitaws.AwsIamUsecase
}

func NewIamService(uc *bitaws.AwsIamUsecase, logger log.Logger) *IamService {
	return &IamService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *IamService) AddAkSecret(ctx context.Context, req *pb.AddAkSecretRequest) (*pb.AddAkSecretReply, error) {
	err := s.uc.AddAkSecret(ctx, domain.NewAddAwsIamFromApi(req))
	if err != nil {
		return nil, err
	}

	return &pb.AddAkSecretReply{}, nil
}
func (s *IamService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleReply, error) {
	return &pb.CreateRoleReply{}, nil
}

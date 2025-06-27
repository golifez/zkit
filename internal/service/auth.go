package service

import (
	"context"

	pb "github.com/golifez/zkit/api/auth/v1"
	"github.com/golifez/zkit/internal/biz"
	"github.com/golifez/zkit/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	log *log.Helper
	uc  *biz.AutherUsecase
}

func NewAuthService(uc *biz.AutherUsecase, logger log.Logger) *AuthService {
	return &AuthService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

// GenToken 生成token
func (s *AuthService) GenToken(ctx context.Context, req *pb.GenTokenRequest) (*pb.GenTokenReply, error) {
	jwt := domain.NewGenTokenRequestFromApi(req)
	token, err := s.uc.GenToken(ctx, jwt)
	if err != nil {
		return nil, err
	}
	return &pb.GenTokenReply{
		Token:   token,
		Message: "生成token成功",
	}, nil
}

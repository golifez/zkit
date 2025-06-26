package service

import (
	"context"

	pb "github.com/golifez/zkit/api/auth/v1"
	"github.com/golifez/zkit/internal/biz"
	"github.com/golifez/zkit/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	log *log.Helper
	uc  *biz.AutherUsecase
}

func NewAuthService(uc *biz.AutherUsecase, logger log.Logger) *AuthService {
	return &AuthService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	auth := domain.Auther{
		Username: req.Username,
		Password: req.Password,
	}
	token, err := s.uc.VerificationAccount(ctx, &auth)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Message: "登录成功",
		Token:   token,
	}, nil
}

// Register 注册
func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	auth := domain.Auther{
		Username: req.Username,
		Password: req.Password,
		NickName: req.Nickname,
	}
	err := s.uc.Register(ctx, &auth)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		Code:    0,
		Message: "注册成功",
	}, nil
}

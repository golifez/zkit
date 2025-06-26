package service

import (
	"context"

	v1 "github.com/golifez/zkit/api/helloworld/v1"
	"github.com/golifez/zkit/internal/biz"
	"github.com/golifez/zkit/internal/errs"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	// errors "github.com/go-kratos/kratos/v2/errors"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	log *log.Helper
	uc  *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	Claims, ok := jwt.FromContext(ctx)
	if ok {
		s.log.Infof("JWT Claims: %+v", Claims)
	}
	if in.Name != "12" {
		// err := errors.New(500, "USER_NAME_EMPTY", "user name is empty")
		return nil, errs.ErrUserNotFound.WithMetadata(map[string]string{
			"foo": "bar",
		})
	}
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

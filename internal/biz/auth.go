package biz

import (
	"context"

	"github.com/golifez/zkit/internal/auth"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

// type Auther struct {
// 	Username string
// 	Passwd   string
// 	NickName string
// }

type AutherRepo interface {
	Login(context.Context, *domain.Auther) (*domain.Auther, error)
	Register(context.Context, *domain.Auther) error
}

// GreeterUsecase is a Greeter usecase.
type AutherUsecase struct {
	repo AutherRepo
	c    *conf.Config
	log  *log.Helper
}

func NewAutherUsecase(repo AutherRepo, c *conf.Config, logger log.Logger) *AutherUsecase {
	return &AutherUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		c:    c}
}

func (a *AutherUsecase) VerificationAccount(ctx context.Context, g *domain.Auther) (token string, err error) {
	g, err = a.repo.Login(ctx, g)
	if err != nil {
		return "", err
	}
	token = auth.GetToken(g.Uid, a.c.Secretkey.Jwtkey)
	if err != nil {
		return "", err
	}
	return token, err
}

// Register 注册
func (a *AutherUsecase) Register(ctx context.Context, g *domain.Auther) (err error) {
	err = a.repo.Register(ctx, g)
	if err != nil {
		return err
	}
	return nil
}

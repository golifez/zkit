package biz

import (
	"context"

	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/domain"
	"github.com/golifez/zkit/utils"

	"github.com/go-kratos/kratos/v2/log"
)

// type Auther struct {
// 	Username string
// 	Passwd   string
// 	NickName string
// }

// type AutherRepo interface {
// }

// GreeterUsecase is a Greeter usecase.
type AutherUsecase struct {
	// repo AutherRepo
	c   *conf.Config
	log *log.Helper
}

func NewAutherUsecase(c *conf.Config, logger log.Logger) *AutherUsecase {
	return &AutherUsecase{
		// repo: repo,
		log: log.NewHelper(logger),
		c:   c}
}

func (a *AutherUsecase) GenToken(ctx context.Context, g *domain.Jwt) (token string, err error) {
	token = utils.NewJwt().GetToken(g.Claims, g.Key)

	// if token != "" {
	// 	return token, err
	// }
	return token, err
}

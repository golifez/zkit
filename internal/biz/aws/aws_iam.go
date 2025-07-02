package aws

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/domain"
)

type AwsIamRepo interface {
	AddAkSecret(context.Context, *domain.AwsIam) error
	AddUser(context.Context, *domain.AwsIam) (*domain.AwsIam, error)
}

type AwsIamUsecase struct {
	repo AwsIamRepo
	c    *conf.Config
	log  *log.Helper
}

func NewAwsIamUsecase(repo AwsIamRepo, c *conf.Config, logger log.Logger) *AwsIamUsecase {
	return &AwsIamUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		c:    c}
}

// 新增AK秘钥绑定
func (a *AwsIamUsecase) AddAkSecret(ctx context.Context, iam *domain.AwsIam) error {
	return a.repo.AddAkSecret(ctx, iam)
}

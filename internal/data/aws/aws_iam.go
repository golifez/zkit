package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golifez/zkit/errs"
	biz "github.com/golifez/zkit/internal/biz/aws"
	"github.com/golifez/zkit/internal/client"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/data/ent/aws_iam"
	"github.com/golifez/zkit/internal/domain"
)

// import (
// 	"context"

// 	"github.com/go-kratos/kratos/v2/log"
// 	biz "github.com/golifez/zkit/internal/biz/aws"
// 	"github.com/golifez/zkit/internal/conf"
// 	"github.com/golifez/zkit/internal/data"
// 	"github.com/golifez/zkit/internal/domain"
// )

type AwsIamRepo struct {
	// data *client.Data
	log  *log.Helper
	c    *conf.Config
	data *client.Data
}

func NewAwsIamRepo(data *client.Data, c *conf.Config, logger log.Logger) biz.AwsIamRepo {
	return &AwsIamRepo{
		data: data,
		log:  log.NewHelper(logger),
		c:    c,
	}
}

func (a *AwsIamRepo) AddAkSecret(ctx context.Context, iam *domain.AwsIam) error {
	// 先查询账号是否存在
	accountnum, err := a.data.DB.Aws_iam.Query().Where(aws_iam.AccountIDEQ(iam.AccountId)).Count(ctx)
	if err != nil {
		return err
	}
	if accountnum > 0 {
		return errs.ErrDataAlreadyExists
	}
	// 不存在则创建
	_, err = a.data.DB.Aws_iam.Create().
		SetAccountID(iam.AccountId).
		SetAccessKey(iam.AccessKey).
		SetSecretKey(iam.SecretKey).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *AwsIamRepo) AddUser(ctx context.Context, iam *domain.AwsIam) (*domain.AwsIam, error) {
	return nil, nil
}

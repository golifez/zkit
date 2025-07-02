package client

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/data/ent"
	"github.com/golifez/zkit/internal/data/ent/migrate"
	"github.com/google/wire"
)

// 移除数据库连接相关函数到 data 层
var ProviderSet = wire.NewSet(NewData, NewEntClient)

// Data .
type Data struct {
	// TODO wrapped database client
	DB *ent.Client
}

// 初始化数据库连接
func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	helper := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))
	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		helper.Fatal("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		helper.Fatal("failed creating schema resources: %v", err)
	}
	return client
}

// NewData .
func NewData(db *ent.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")

	}
	return &Data{
		DB: db,
	}, cleanup, nil
}

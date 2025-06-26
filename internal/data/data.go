package data

import (
	"context"

	"github.com/golifez/zkit/internal/conf"

	"github.com/golifez/zkit/internal/data/ent"
	"github.com/golifez/zkit/internal/data/ent/migrate"

	// init mysql driver
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClinet, NewGreeterRepo, NewAutherRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *ent.Client
}

// 初始化数据库连接
func NewEntClinet(conf *conf.Data, logger log.Logger) *ent.Client {
	helper := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))
	clinet, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		helper.Fatal("failed opening connection to sqlite: %v", err)
	}
	if err := clinet.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		helper.Fatal("failed creating schema resources: %v", err)
	}
	return clinet
}

// NewData .
func NewData(db *ent.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")

	}
	return &Data{
		db: db,
	}, cleanup, nil
}

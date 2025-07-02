package data

import (
	_ "github.com/go-sql-driver/mysql"
	dataaws "github.com/golifez/zkit/internal/data/aws"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	// // 从 client 层移回的数据库连接相关函数\
	// NewEntClient,
	// NewData,
	dataaws.NewAwsIamRepo,
)

// // Data .
// type Data struct {
// 	DB *ent.Client
// }

// // // 初始化数据库连接
// func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
// 	helper := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))
// 	client, err := ent.Open(
// 		conf.Database.Driver,
// 		conf.Database.Source,
// 	)
// 	if err != nil {
// 		helper.Fatal("failed opening connection to sqlite: %v", err)
// 	}
// 	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
// 		helper.Fatal("failed creating schema resources: %v", err)
// 	}
// 	return client
// }

// // // NewData .
// func NewData(db *ent.Client, logger log.Logger) (*Data, func(), error) {
// 	cleanup := func() {
// 		log.NewHelper(logger).Info("closing the data resources")
// 	}
// 	return &Data{
// 		DB: db,
// 	}, cleanup, nil
// }

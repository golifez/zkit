package data

import (

	// init mysql driver

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
// var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	// TODO wrapped database client
	// db *ent.Client
}

// 初始化数据库连接
// func NewEntClinet(conf *conf.Data, logger log.Logger) *ent.Client {
// 	helper := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))
// 	clinet, err := ent.Open(
// 		conf.Database.Driver,
// 		conf.Database.Source,
// 	)
// 	if err != nil {
// 		helper.Fatal("failed opening connection to sqlite: %v", err)
// 	}
// 	if err := clinet.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
// 		helper.Fatal("failed creating schema resources: %v", err)
// 	}
// 	return clinet
// }

// NewData .
// func NewData(logger log.Logger) (*Data, func(), error) {
// 	cleanup := func() {
// 		log.NewHelper(logger).Info("closing the data resources")

// 	}
// 	return &Data{
// 		// db: db,
// 	}, cleanup, nil
// }

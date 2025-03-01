package mysql

import (
	"douyin_mall/cart/biz/model"
	"douyin_mall/cart/conf"
	"gorm.io/plugin/opentelemetry/tracing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := conf.GetConf().MySQL.DSN
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:    true,
			TranslateError: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.CartItem{})
	if err != nil {
		panic(err)
	}
	if conf.GetConf().Env == "dev" {
		DB = DB.Debug()
	}
}

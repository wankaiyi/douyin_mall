package mysql

import (
	"douyin_mall/doubao_ai/biz/model"
	"douyin_mall/doubao_ai/conf"
	"gorm.io/plugin/opentelemetry/tracing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	DB.DisableForeignKeyConstraintWhenMigrating = true
	err = DB.AutoMigrate(&model.Message{})
	if err != nil {
		panic(err)
	}
	if conf.GetConf().Env == "dev" {
		DB = DB.Debug()
	}
}

package mysql

import (
	"douyin_mall/user/biz/model"
	"douyin_mall/user/conf"
	"gorm.io/driver/mysql"
	"gorm.io/plugin/opentelemetry/tracing"

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
	DB.DisableForeignKeyConstraintWhenMigrating = true
	err = DB.AutoMigrate(&model.User{}, &model.Address{})
	if err != nil {
		panic(err)
	}
	if conf.GetConf().Env == "dev" {
		DB = DB.Debug()
	}
}

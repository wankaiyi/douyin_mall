package casbin

import (
	"douyin_mall/auth/conf"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Enforcer *casbin.Enforcer

func Init() {
	dsn := conf.GetConf().Mysql.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		klog.Errorf("连接数据库失败: %v", err)
		panic(err)
	}

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		klog.Errorf("创建casbin适配器失败: %v", err)
		panic(err)
	}

	Enforcer, err = casbin.NewEnforcer(conf.GetConf().Casbin.ModelPath, adapter)
	if err != nil {
		klog.Errorf("创建casbin执行器失败: %v", err)
		panic(err)
	}

}

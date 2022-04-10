package bootstrap

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"peng-api/global"
)

func InitializeCasbin() *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(global.App.DB)
	if err != nil {
		global.App.Log.Error("Adapter create failed,err:", zap.Any("err", err))
	}
	enforcer, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		global.App.Log.Error("NewEnforcer failed,err:", zap.Any("err", err))
	}
	return enforcer
}

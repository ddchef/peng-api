package bootstrap

import (
	"peng-api/app/models"
	"peng-api/global"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

func InitializeCasbin() *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(global.App.DB, &models.CasbinRule{})
	if err != nil {
		global.App.Log.Error("Adapter create failed,err:", zap.Any("err", err))
	}
	enforcer, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		global.App.Log.Error("NewEnforcer failed,err:", zap.Any("err", err))
	}
	return enforcer
}

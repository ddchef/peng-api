package global

import (
	"peng-api/config"

	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
	Casbin      *casbin.Enforcer
}

var App = new(Application)

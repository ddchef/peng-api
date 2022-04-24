package main

import (
	"fmt"
	"peng-api/bootstrap"
	"peng-api/global"
)

// @title Peng API
// @version 1.0
// @description example
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 初始化 Casbin
	global.App.Casbin = bootstrap.InitializeCasbin()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	// 初始化验证器
	if err := bootstrap.InitializeValidator("zh"); err != nil {
		fmt.Print("初始化验证器失败")
	}
	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()
	// 启动服务
	bootstrap.RunServer()
}

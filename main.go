package main

import (
	"boost/data/server/core"
	"boost/data/server/global"
	"boost/data/server/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

// @BasePath
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm_A() // gorm连接数据库
	//global.GVA_DB_B = initialize.Gorm_B() // gorm连接数据库
	//global.GVA_DB_C = initialize.Gorm_C() // gorm连接数据库
	initialize.Timer()

	if global.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}

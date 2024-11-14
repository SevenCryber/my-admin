package main

import (
	"github.com/SevenCryber/my-admin/server/core"
	"github.com/SevenCryber/my-admin/server/global"
	"github.com/SevenCryber/my-admin/server/initialize"
	"go.uber.org/zap"
)

func main() {
	//初始化Viper,读取配置信息
	global.YAC_VP = core.Viper()
	initialize.OtherInit()
	// 初始化zap日志库
	global.YAC_LOG = core.Zap()
	zap.ReplaceGlobals(global.YAC_LOG)
	// gorm连接数据库
	global.YAC_DB = initialize.Gorm()
	//初始化定时任务
	initialize.Timer()
	// 初始化表
	if global.YAC_DB != nil {
		initialize.RegisterTables()
		// 程序结束前关闭数据库链接
		db, _ := global.YAC_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}

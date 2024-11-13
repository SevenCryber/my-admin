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
}

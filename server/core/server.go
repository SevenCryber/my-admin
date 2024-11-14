package core

import (
	"fmt"
	"github.com/SevenCryber/my-admin/server/global"
	"github.com/SevenCryber/my-admin/server/initialize"
	"github.com/SevenCryber/my-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.YAC_CONFIG.System.UseMultipoint || global.YAC_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		initialize.RedisList()
	}

	// 从db加载jwt数据
	if global.YAC_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.YAC_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.YAC_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`欢迎使用`, address)
	global.YAC_LOG.Error(s.ListenAndServe().Error())
}

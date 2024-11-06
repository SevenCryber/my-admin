package main

import (
	"github.com/SevenCryber/my-admin/server/core"
	"github.com/SevenCryber/my-admin/server/global"
)

func main() {
	//初始化Viper
	global.YAC_VP = core.Viper()
}

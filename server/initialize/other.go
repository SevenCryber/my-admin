package initialize

import (
	"github.com/SevenCryber/my-admin/server/global"
	"github.com/SevenCryber/my-admin/server/utils"
	"github.com/SevenCryber/my-admin/server/utils/cache"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.YAC_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.YAC_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = cache.NewCache(
		cache.SetDefaultExpire(dr),
	)
}

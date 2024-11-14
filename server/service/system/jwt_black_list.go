package system

import (
	"github.com/SevenCryber/my-admin/server/global"
	"github.com/SevenCryber/my-admin/server/model/system"
	"go.uber.org/zap"
)

func LoadAll() {
	var data []string
	err := global.YAC_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.YAC_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}

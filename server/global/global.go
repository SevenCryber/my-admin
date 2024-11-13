package global

import (
	"github.com/SevenCryber/my-admin/server/config"
	"github.com/SevenCryber/my-admin/server/utils/cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	YAC_VP     *viper.Viper
	YAC_CONFIG config.Server
	BlackCache cache.Cache
	YAC_LOG    *zap.Logger
)

package global

import (
	"github.com/SevenCryber/my-admin/server/config"
	"github.com/SevenCryber/my-admin/server/utils/cache"
	"github.com/SevenCryber/my-admin/server/utils/timer"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	YAC_VP            *viper.Viper
	YAC_CONFIG        config.Server
	BlackCache        cache.Cache
	YAC_LOG           *zap.Logger
	YAC_ACTIVE_DBNAME *string
	YAC_DB            *gorm.DB
	YAC_Timer         timer.Timer = timer.NewTimerTask()
	YAC_REDIS         redis.UniversalClient
	YAC_REDISList     map[string]redis.UniversalClient
)

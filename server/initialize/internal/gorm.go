package internal

import (
	"github.com/SevenCryber/my-admin/server/config"
	"github.com/SevenCryber/my-admin/server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type _gorm struct{}

var Gorm = new(_gorm)

// Config gorm 自定义配置
// Author SevenCryber
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	var general config.GeneralDB
	switch global.YAC_CONFIG.System.DbType {
	case "mysql":
		general = global.YAC_CONFIG.Mysql.GeneralDB
	default:
		general = global.YAC_CONFIG.Mysql.GeneralDB
	}
	return &gorm.Config{
		Logger: logger.New(NewWriter(general, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

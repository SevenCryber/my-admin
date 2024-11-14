package initialize

import (
	"github.com/SevenCryber/my-admin/server/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.YAC_CONFIG.System.DbType {
	case "mysql":
		global.YAC_ACTIVE_DBNAME = &global.YAC_CONFIG.Mysql.Dbname
		return GormMysql()
	default:
		global.YAC_ACTIVE_DBNAME = &global.YAC_CONFIG.Mysql.Dbname
		return GormMysql()
	}
}

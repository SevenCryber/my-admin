package initialize

import (
	"github.com/SevenCryber/my-admin/server/global"
	"github.com/SevenCryber/my-admin/server/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
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

func RegisterTables() {
	db := global.YAC_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},
		system.SysParams{},
	)
	if err != nil {
		global.YAC_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel()

	if err != nil {
		global.YAC_LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.YAC_LOG.Info("register table success")
}

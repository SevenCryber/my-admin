package initialize

import "github.com/SevenCryber/my-admin/server/global"

func bizModel() error {
	db := global.YAC_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}

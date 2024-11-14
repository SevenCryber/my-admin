package system

import "github.com/SevenCryber/my-admin/server/global"

type JwtBlacklist struct {
	global.YAC_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

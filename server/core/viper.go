package core

import (
	"flag"
	"github.com/SevenCryber/my-admin/server/core/internal"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "config file path")
		flag.Parse()
		// 判断命令行参数是否为空
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {

				}
			}
		}
	}
}

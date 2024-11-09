package core

import (
	"flag"
	"fmt"
	"github.com/SevenCryber/my-admin/server/core/internal"
	"github.com/SevenCryber/my-admin/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var config string
	//先判断传入的参数，如果没有就从命令行读取
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "config file path")
		flag.Parse()
		// 判断命令行参数是否为空，为空就从环境变量中读取
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
				case gin.TestMode:
					config = internal.ConfigTestFile
				}
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config)
			} else {
				config = configEnv
				fmt.Printf("正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("正在使用func Viper()传递的值,config的路径为%s\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.YAC_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.YAC_CONFIG); err != nil {
		panic(err)
	}
	return v
}

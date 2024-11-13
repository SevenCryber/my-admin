package core

import (
	"fmt"
	"github.com/SevenCryber/my-admin/server/core/internal"
	"github.com/SevenCryber/my-admin/server/global"
	"github.com/SevenCryber/my-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.YAC_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.YAC_CONFIG.Zap.Director)
		_ = os.Mkdir(global.YAC_CONFIG.Zap.Director, os.ModePerm)
	}
	levels := global.YAC_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.YAC_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

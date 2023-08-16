package initialize

import (
	"HiChat/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	v := viper.New()
	configFile := "../HiChat/config-debug.yaml"

	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServiceConfig); err != nil {
		panic(err)
	}
	zap.S().Info("配置信息", global.ServiceConfig)
}

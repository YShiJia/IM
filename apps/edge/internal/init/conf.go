package init

import (
	"fmt"
	"github.com/YShiJia/IM/apps/edge/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.AddConfigPath("./apps/edge/etc") // 配置文件的路径，可以设置多个路径
	viper.SetConfigName("edge")            // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")            // 设置配置文件类型为yaml

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("reading config file failed:  %v", err)
	}

	if err := viper.Unmarshal(&config.Conf); err != nil {
		return fmt.Errorf("unmarshaling config failed, %v", err)
	}
	log.Infof("init config success, config: %#v", config.Conf)
	return nil
}

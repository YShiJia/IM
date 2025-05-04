package init

import (
	"fmt"
	"github.com/YShiJia/IM/apps/edge/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func InitConfig(configPath string) error {

	if configPath == "" {
		setDefaultConfigFile()
	} else {
		setSpecialConfigFile(configPath)
	}

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

func setDefaultConfigFile() {
	viper.AddConfigPath("./apps/edge/etc") // 配置文件的路径，可以设置多个路径
	viper.SetConfigName("edge")            // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")            // 设置配置文件类型为yaml
}

func setSpecialConfigFile(configPath string) {
	// 取目录路径
	configDir := filepath.Dir(configPath)
	// 取文件名
	fileName := filepath.Base(configPath)
	// 取扩展名（带点的，比如 ".yaml"）
	ext := filepath.Ext(fileName)
	// 去掉扩展名，得到纯文件名
	configName := strings.TrimSuffix(fileName, ext)
	// 把 .yaml -> yaml
	configType := strings.TrimPrefix(ext, ".")

	viper.AddConfigPath(configDir)  // 配置文件的路径，可以设置多个路径
	viper.SetConfigName(configName) // 配置文件名（不带扩展名）
	viper.SetConfigType(configType) // 设置配置文件类型为yaml
}

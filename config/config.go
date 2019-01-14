package config

import (
	"github.com/go-ini/ini"
)

var (
	AppConfig 	= &appConfig{}
	
)

//加载基础配置
func LoadSystemConfig(iniPath string) (msg string, err error) {
	configSetUp, err := ini.Load(iniPath)
	if err != nil {
		return "app基础配置加载失败！", err
	}
	err = configSetUp.Section("App").MapTo(AppConfig)
	if err != nil {
		return "app基础配置加载失败！", err
	}
	return "app基础配置全部加载完成！", nil
}
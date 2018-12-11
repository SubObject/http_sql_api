package config

import (
	"github.com/go-ini/ini"
)

var (
	AppConfig 	= &appConfig{}
	MysqlConfig = &mysqlConfig{}
	RedisConfig = &redisConfig{}
)

//加载配置
func LoadSystemConfig(iniPath string) (msg string, err error) {
	configSetUp, err := ini.Load(iniPath)
	if err != nil {
		return "配置加载失败！", err
	}
	err = configSetUp.Section("App").MapTo(AppConfig)
	if err != nil {
		return "app基础配置加载失败！", err
	}
	err = configSetUp.Section("Mysql").MapTo(MysqlConfig)
	if err != nil {
		return "Mysql数据库配置加载失败！", err
	}
	err = configSetUp.Section("Redis").MapTo(RedisConfig)
	if err != nil {
		return "Redis配置加载失败！", err
	}
	return "配置全部加载完成！", nil
}

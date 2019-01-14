package config

import (
	"github.com/go-ini/ini"
)

var (
	RedisConfig = &redisConfig{}
)
//加载Redis配置
func LoadRedisConfig(redisPath string) (msg string,err error){
	redisSetUp, err := ini.Load(redisPath)
	if err != nil {
		return "Redis配置加载失败！", err
	}
	err = redisSetUp.Section("Mysql").MapTo(MysqlConfig)
	if err != nil {
		return "Redis配置加载失败！", err
	}
	return "Redis配置全部加载完成！", nil
}
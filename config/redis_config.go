package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	RedisConfig = &redisConfig{}
)
//加载Redis配置
func LoadRedisConfig(redisPath string) (err error){
	fmt.Println("开始加载Redis基础配置...")
	redisSetUp, err := ini.Load(redisPath)
	if err != nil {
		fmt.Println("Redis配置加载失败！")
		return err
	}
	err = redisSetUp.Section("Mysql").MapTo(MysqlConfig)
	if err != nil {
		fmt.Println("Redis配置加载失败！")
		return err
	}
	fmt.Println("Redis配置全部加载完成！")
	return nil
}
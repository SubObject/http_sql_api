package config

import (
	"github.com/go-ini/ini"
)

var (
	MysqlConfig = &mysqlConfig{}
)
//加载数据库配置
func LoadDataBaseConfig(dataPath string) (msg string, err error) {
	dataBaseSetUp, err := ini.Load(dataPath)
	if err != nil {
		return "数据库配置加载失败！", err
	}
	err = dataBaseSetUp.Section("Mysql").MapTo(MysqlConfig)
	if err != nil {
		return "数据库配置加载失败！", err
	}
	return "数据库配置全部加载完成！", nil
}
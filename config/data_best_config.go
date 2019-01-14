package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	MysqlConfig = &mysqlConfig{}
)
//加载数据库配置
func LoadDataBaseConfig(dataPath string) (err error) {
	fmt.Println("开始加载Mysql基础配置...")
	dataBaseSetUp, err := ini.Load(dataPath)
	if err != nil {
		fmt.Println("数据库配置加载失败！")
		return err
	}
	err = dataBaseSetUp.Section("Mysql").MapTo(MysqlConfig)
	if err != nil {
		fmt.Println("数据库配置加载失败！")
		return err
	}
	fmt.Println("数据库配置全部加载完成！")
	return nil
}
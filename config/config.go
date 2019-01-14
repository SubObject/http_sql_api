package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	AppConfig 	= &appConfig{}
	
)

//加载基础配置
func LoadSystemConfig(iniPath string) (err error) {
	fmt.Println("启动程序,开始加载基础配置...")
	configSetUp, err := ini.Load(iniPath)
	if err != nil {
		fmt.Println("app基础配置加载失败！")
		return err
	}
	err = configSetUp.Section("App").MapTo(AppConfig)
	if err != nil {
		fmt.Println("app基础配置加载失败！")
		return err
	}
	fmt.Println("app基础配置全部加载完成！")
	return nil
}
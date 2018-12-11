package main

import (
	"fmt"
	"flag"
	"os"

	"http_sql_api/config"
)

var (
	systemConfigFiles = flag.String("configFile","conf/appConfig.ini","App配置文件")
)

func main() {
	fmt.Println("启动程序,开始加载基础配置...")
	flag.Parse()
	
	configSetUp, err := config.LoadSystemConfig(*systemConfigFiles)
	fmt.Println(configSetUp)
	if err != nil {
		os.Exit(3)
	}
	
}
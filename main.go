package main

import (
	"fmt"
	"flag"
	"os"

	"http_sql_api/config"
	"http_sql_api/service"
)

var (
	systemConfigFiles = flag.String("configFile","conf/appConfig.ini","App配置文件")
	dataBestConfigFiles = flag.String("mysqlconfigFile","conf/mysqlConfig.ini","Mysql配置文件")
	redisConfigFiles = flag.String("redisconfigFile","conf/redisConfig.ini","Redis配置文件")
)

func main() {
	fmt.Println("启动程序,开始加载基础配置...")
	flag.Parse()
	
	configSetUp, err := config.LoadSystemConfig(*systemConfigFiles)
	fmt.Println(configSetUp)
	if err != nil {
		os.Exit(3)
	}
	fmt.Println("开始加载Mysql基础配置...")
	mysqlSetUp, mysql_err :=  config.LoadDataBaseConfig(*dataBestConfigFiles)
	fmt.Println(mysqlSetUp)
	if mysql_err != nil {
		os.Exit(3)
	}
	fmt.Println("开始加载Redis基础配置...")
	redisSetUp, redis_err :=  config.LoadRedisConfig(*redisConfigFiles)
	fmt.Println(redisSetUp)
	if redis_err != nil {
		os.Exit(3)
	}
	server := service.NewServer()
	server.Run(config.AppConfig.ListenPort)
}
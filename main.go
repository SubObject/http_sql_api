package main

import (
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
	
	flag.Parse()
	//加载基础配置
	err := config.LoadSystemConfig(*systemConfigFiles)
	if err != nil {
		os.Exit(3)
	}
	//加载Mysql基础配置
	mysql_err :=  config.LoadDataBaseConfig(*dataBestConfigFiles)
	if mysql_err != nil {
		os.Exit(3)
	}
	//加载Redis基础配置
	redis_err :=  config.LoadRedisConfig(*redisConfigFiles)
	if redis_err != nil {
		os.Exit(3)
	}
	server := service.NewServer()
	server.Run(config.AppConfig.ListenPort)
}
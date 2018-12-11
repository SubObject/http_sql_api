package config

//App基础配置
type appConfig struct {
	ListenPort				string			//监听端口
	AppKey					string			//App加密密钥
	DataBaseType			string			//采用数据库类型
	ReadWriteSeparation		bool			//读写分离
}

//数据库配置
type mysqlConfig struct {
	DB_Host					string			//数据库地址
	DB_User					string			//数据库用户名
	DB_Pwds					string			//数据库密码
	DB_Port					int				//数据库端口
	DB_Name					string			//数据库名称
	Master_Slave			int				//1：主库；0：从库
}

//Redis配置
type redisConfig struct {
	Redis_Host				string			//redis地址
	Redis_Port				int				//redis端口
	Redis_pwds				string			//redis密码
}
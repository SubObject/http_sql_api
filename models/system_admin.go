package models

type SystemAdmin struct {
	Id			int64		`json:"id"`			//ID
	UserName	string		`json:"username"`	//用户名
	Pwd			string		`json:"pwd"`		//密码
	Creade		int64		`json:creade`		//创建时间
	UpDate		int64		`json:"update"`		//更新时间
	FullName	string		`json:"fullname"`	//姓名
}
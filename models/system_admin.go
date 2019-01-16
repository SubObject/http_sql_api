package models

type SystemAdmin struct {
	Id			int64		`json:"id" bson:"id"`			//ID
	UserName	string		`json:"username" bson:"username"`	//用户名
	Pwd			string		`json:"pwd" bson:"pwd"`		//密码
	Creade		int64		`json:"creade" bson:"creade"`		//创建时间
	UpDate		int64		`json:"update" bson:"update"`		//更新时间
	FullName	string		`json:"fullname" bson:"fullname"`	//姓名
}
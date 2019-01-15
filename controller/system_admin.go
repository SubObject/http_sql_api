package controller

import (
	//"fmt"
	"net/http"
	"time"

	"http_sql_api/outputformat"
	"http_sql_api/sql_curd"
	"http_sql_api/models"
)
var sql_model sql_curd.Models
func SelectSystemAdmin()  http.HandlerFunc  {
	prt := outputformat.JsonOut{}

	
	username := "herenshan"
	Whe_ary := sql_curd.SetMapOut()
	Whe_ary["username"]=sql_curd.Setwhere{"=",username}

	ord_ary :=	sql_curd.SetMapOut()
	ord_ary["id"]="desc"
	ord_ary["creade"]="desc"

	
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sq").Field("*").Where(Whe_ary).OrderBy(ord_ary).Limit(1,20).Select()

	data := outputformat.MapOut()	
	
	data["list"]=list

	
	prt.Code = 200
	prt.Msg = "通讯成功！"
	prt.Data=data
	prt.ErrMsg=err
	
	return outputformat.OutPutJson(prt)
}
//写入管理员
func AddSystemAdmin() http.HandlerFunc {
	prt := outputformat.JsonOut{}
	var userModel models.SystemAdmin

	gentor1, err := outputformat.NewIDGenerator().SetWorkerId(100).Init()
	gid, err := gentor1.NextId()
	userModel.Id = gid
	userModel.UserName = "herenshan110"
	userModel.Pwd = outputformat.Md5("123456")
	userModel.Creade = time.Now().Unix()
	userModel.UpDate = time.Now().Unix()
	userModel.FullName = "何仁山"

	cont, err := sql_model.Db().Save(&userModel)

	data := outputformat.MapOut()	
	data["list"]=cont
	prt.Code = 200
	prt.Msg = "通讯成功！"
	prt.Data=data
	prt.ErrMsg=err
	return outputformat.OutPutJson(prt)
}


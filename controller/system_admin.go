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
func selectSystemAdmin(req *http.Request)  (prt outputformat.JsonOut)  {

	// username := "herenshan"
	// Whe_ary := sql_curd.SetMapOut()
	// Whe_ary["username"]=sql_curd.Setwhere{"like",username}

	ord_ary :=	sql_curd.SetMapOut()
	ord_ary["id"]="desc"
	ord_ary["creade"]="desc"

	
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sq").Field("*").OrderBy(ord_ary).Limit(0,20).Select()

	data := outputformat.MapOut()	
	
	data["list"]=list

	
	prt.Code = 200
	prt.Msg = "通讯成功！"
	prt.Data=data
	prt.ErrMsg=err
	
	return prt
}
//写入管理员
func addAdminHandle(req *http.Request)  (prt outputformat.JsonOut)  {
	//prt := outputformat.JsonOut{}
	var userModel models.SystemAdmin

	username := req.FormValue("username")
	pwd := req.FormValue("pwd")
	qrpwd := req.FormValue("qrpwd")
	fullname := req.FormValue("fullname")
	if username == "" {
		prt.Code=1002
		prt.Msg="对不起，请输入用户名！"
		return prt
	}
	if pwd == "" {
		prt.Code=1003
		prt.Msg="对不起，请输入密码！"
		return prt
	}
	if qrpwd == "" {
		prt.Code=1004
		prt.Msg="对不起，请确认密码输入！"
		return prt
	}
	if fullname == "" {
		prt.Code=1005
		prt.Msg="对不起，请输入管理员姓名！"
		return prt
	}
	if pwd != qrpwd {
		prt.Code=1006
		prt.Msg="对不起，您两次输入的密码不一致！"
		return prt
	}
	// gentor1, err := outputformat.NewIDGenerator().SetWorkerId(100).Init()
	// gid, err := gentor1.NextId()
	// userModel.Id = gid
	userModel.UserName = username
	userModel.Pwd = outputformat.Md5("123456")
	userModel.Creade = time.Now().Unix()
	userModel.UpDate = time.Now().Unix()
	userModel.FullName = "何仁山"

	// cont, err := sql_model.Db().Save(&userModel)
	cont, err := sql_model.Db().Insert(&userModel)

	data := outputformat.MapOut()	
	data["list"]=cont
	prt.Code = 200
	prt.Msg = "通讯成功！"
	prt.Data=data
	prt.ErrMsg=err
	return prt
}


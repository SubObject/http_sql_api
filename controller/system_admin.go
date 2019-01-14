package controller

import (
	//"fmt"
	"net/http"

	"http_sql_api/outputformat"
	"http_sql_api/sql_curd"
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



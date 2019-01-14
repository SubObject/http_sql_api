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

	
	// username := "herenshan"
	// Whe_ary := sql_curd.SetMapOut()
	// Whe_ary["username"]=sql_curd.Setwhere{"=",username}

	ord_ary :=	sql_curd.SetMapOut()
	ord_ary["id"]="desc"
	ord_ary["creader"]="desc"

	//list,_ :=sql_model.Db().TableNames("system_admin").Where(Whe_ary)->Find()
	// list := sql_model.Db()

	// fmt.Println(list)

	prt.Code = 200
	prt.Msg = "通讯成功！=>SelectSystemAdmin"
	prt.Data=ord_ary
	
	return outputformat.OutPutJson(prt)
}



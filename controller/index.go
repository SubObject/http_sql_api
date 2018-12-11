package controller

import (
	"net/http"

	"http_sql_api/outputformat"
)

func Index() http.HandlerFunc {
	prt := outputformat.JsonOut{}
	prt.Code = 200
	prt.Msg = "通讯成功！"
	return outputformat.OutPutJson(prt)
}
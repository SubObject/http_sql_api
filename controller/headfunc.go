package controller

import (
	"net/http"

	"github.com/unrolled/render"
)
//获取查询列表
func SelectSystemAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		formatter := render.New(render.Options{IndentJSON:true})
		formatter.JSON(w,http.StatusOK,selectSystemAdmin(req))
	}
}
//写入会员
func AddSystemAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		formatter := render.New(render.Options{IndentJSON:true})
		formatter.JSON(w,http.StatusOK,addAdminHandle(req))
	}
}
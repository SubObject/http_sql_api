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

func AddSystemAdmining() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		formatter := render.New(render.Options{IndentJSON:true})
		formatter.JSON(w,http.StatusOK,addAdminHandleing(req))
	}
}

//更新数据
func UpdateAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		formatter := render.New(render.Options{IndentJSON:true})
		formatter.JSON(w,http.StatusOK,updateAdminRun(req))
	}
}

//删除数据
func DelDataAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		formatter := render.New(render.Options{IndentJSON:true})
		formatter.JSON(w,http.StatusOK,delSystemAdmin(req))
	}
}
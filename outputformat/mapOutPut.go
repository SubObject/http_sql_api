package outputformat

import (
	"net/http"
	"github.com/unrolled/render"
)
func MapOut() (data map[string]interface{}) {
	data= make(map[string]interface{}) //必可不少，分配内存
	return data
}

func OutPutJson(out_val JsonOut) http.HandlerFunc {
	return func(w http.ResponseWriter,req *http.Request){
		formatter := render.New(render.Options{IndentJSON:true})
		formatter.JSON(w,http.StatusOK,out_val)
	}
}
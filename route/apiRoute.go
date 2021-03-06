package route

import (
	"http_sql_api/controller"

	"github.com/gorilla/mux"
)

//路由
func InitRoute(m *mux.Router) {
	m.HandleFunc("/",controller.Index()).Methods("GET")
	m.HandleFunc("/system_admin",controller.SelectSystemAdmin()).Methods("GET")
	m.HandleFunc("/add_admin",controller.AddSystemAdmin()).Methods("POST")
	m.HandleFunc("/add_admin_ing",controller.AddSystemAdmining()).Methods("POST")
	m.HandleFunc("/edit_admin_ing",controller.UpdateAdmin()).Methods("POST")
	m.HandleFunc("/del_admin_ing",controller.DelDataAdmin()).Methods("POST")
}
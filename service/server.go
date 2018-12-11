package service

import (
	"http_sql_api/route"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	m := mux.NewRouter()
	route.InitRoute(m)
	n.UseHandler(m)
	return n
}
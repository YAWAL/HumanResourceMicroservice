package main

import (
	"github.com/YAWAL/ERP/back-end/HumanResourcesService/router"
	"github.com/YAWAL/ERP/back-end/tool/logging"
	"net/http"
)

func main() {

	r := router.Router

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	logging.Log.Println(srv.ListenAndServe())

}

package main

import (
	"net/http"

	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"github.com/YAWAL/HumanResourceMicroservice/src/router"
)

func main() {

	r := router.Router

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	logging.Log.Println(srv.ListenAndServe())

}

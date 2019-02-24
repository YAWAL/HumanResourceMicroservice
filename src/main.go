package main

import (
	"fmt"
	"net/http"

	"github.com/YAWAL/HumanResourceMicroservice/src/router"
)

func main() {

	r := router.InitRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	fmt.Println(srv.ListenAndServe())

}

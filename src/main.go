package main

import (
	"fmt"
	"github.com/YAWAL/HumanResourceMicroservice/src/config"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"net/http"
	"os"

	"github.com/YAWAL/HumanResourceMicroservice/src/router"
)

func main() {
	conf, err := config.LoadConfig(os.Args[0])
	if err != nil {
		logging.Log.Errorf("Cannot load config: ", err)
	}

	r := router.InitRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    conf.Host,
	}
	fmt.Println(srv.ListenAndServe())

}

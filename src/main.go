package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/YAWAL/HumanResourceMicroservice/src/app"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
)

func main() {

	file := flag.String("config", "../config.json", "config file path")
	flag.Parse()
	srv, err := app.LoadApp(*file)
	if err != nil {
		logging.Log.Println("error during starting app: %v", err)
	}

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	go app.GracefullShutdown(srv, quit, done)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logging.Log.Println("Could not listen on %s: %v\n", os.Args[0], err)
	}
	<-done
	logging.Log.Println("Server stopped")
}

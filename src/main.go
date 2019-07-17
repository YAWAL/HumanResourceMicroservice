package main

import (
	"github.com/YAWAL/HumanResourceMicroservice/src/app"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"net/http"
	"os"
)

func main() {

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	srv, err := app.LoadApp()
	if err != nil {
		logging.Log.Println("error during starting app: %v", err)
	}

	go app.GracefullShutdown(srv, quit, done)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logging.Log.Println("Could not listen on %s: %v\n", os.Args[0], err)
	}
	<-done
	logging.Log.Println("Server stopped")
}

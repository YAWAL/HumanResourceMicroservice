package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/YAWAL/HumanResourceMicroservice/src/repository"

	"github.com/YAWAL/HumanResourceMicroservice/src/config"
	"github.com/YAWAL/HumanResourceMicroservice/src/database"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"github.com/YAWAL/HumanResourceMicroservice/src/router"
)

func LoadApp(file string) (srv *http.Server, err error) {
	// read, config
	conf, err := config.ReadConfig(file)
	if err != nil {
		logging.Log.Errorf("Cannot load config: ", err.Error())
		return nil, err
	}
	// establish connections to DB
	db, err := database.PGconn(conf.Database)
	if err != nil {
		logging.Log.Errorf("Cannot connect to DB: %s", err.Error())
		return nil, err
	}

	//init repos
	repo := repository.NewPostgresRepository(db)

	// init routers
	r := router.InitRouter(repo)

	srv = &http.Server{
		Handler:      r,
		Addr:         conf.Host,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	logging.Log.Infof("Application is running on %s ", conf.Host)

	return srv, nil
}

func GracefullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logging.Log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logging.Log.Println("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

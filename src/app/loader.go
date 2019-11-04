package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/YAWAL/HumanResourceMicroservice/src/config"
	"github.com/YAWAL/HumanResourceMicroservice/src/database"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"github.com/YAWAL/HumanResourceMicroservice/src/repository"
	"github.com/YAWAL/HumanResourceMicroservice/src/router"
)

// LoadApp performs reading config file, creating services and http server
func LoadApp(file string) (srv *http.Server, err error) {
	// read, config
	conf, err := config.ReadConfig(file)
	if err != nil {
		logging.Log.Errorf("Cannot load config: %s", err.Error())
		return nil, err
	}
	// establish connections to MongoDB
	cl, err := database.Mongo(conf.Database)
	if err != nil {
		logging.Log.Errorf("Cannot create MongoDB client: %s", err.Error())
		return nil, err
	}
	logging.Log.Infof("Creating MongoBD client")

	// checking the connection
	err = cl.Ping(context.TODO(), nil)
	if err != nil {
		logging.Log.Errorf("Cannot connect to MongoDB: %s", err.Error())
		return nil, err
	}
	logging.Log.Infof("Connected to MongoBD")

	//init repos
	repo := repository.NewMongoRepository(cl)
	logging.Log.Infof("Initializing MongoBD repository")

	// init routers
	r := router.InitRouter(repo)
	logging.Log.Infof("Initializing routers")

	srv = &http.Server{
		Handler:      r,
		Addr:         conf.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	logging.Log.Infof("Application is running on %s", conf.Port)

	return srv, nil
}

// GracefullShutdown performs grecefull shutdown of server and loggs to stdout
// info about operation success
func GracefullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logging.Log.Info("Server is shutting down...")

	// TODO err = client.Disconnect(context.TODO()) -> https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logging.Log.Errorf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

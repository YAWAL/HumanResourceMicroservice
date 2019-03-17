package app

import (
	"fmt"
	"github.com/YAWAL/HumanResourceMicroservice/src/HRrepository"
	"github.com/YAWAL/HumanResourceMicroservice/src/config"
	"github.com/YAWAL/HumanResourceMicroservice/src/database"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"github.com/YAWAL/HumanResourceMicroservice/src/router"
	"net/http"
	"os"
)

func LoadApp() error {
	// read, config
	conf, err := config.LoadConfig(os.Args[0])
	if err != nil {
		logging.Log.Errorf("Cannot load config: ", err)
		return err
	}
	// establish connections to DB
	databaseConnection, err := database.PGconn(conf.Database)
	if err != nil {
		logging.Log.Errorf("Cannot load config: ", err)
		return err
	}

	//init repos
	hrRepo := HRrepository.NewPostgresRepository(databaseConnection)

	// init handlers

	// init routers
	r := router.InitRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    conf.Host,
	}
	fmt.Println(srv.ListenAndServe())

	return nil
}

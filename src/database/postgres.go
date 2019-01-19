package database

import (
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func PGconn() (db *gorm.DB, err error) {
	// real configs -> config file
	db, err = gorm.Open("postgres", "user=postgres dbname=plant sslmode=disable password=postgres")
	if err != nil {

		logging.Log.Errorf("Error during connection to Postgres has occurred %s", err.Error())
	} else {
		logging.Log.Info("Connection to Pg has been established")
	}
	return db, err
}

package database

import (
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func PGconn(conf Config) (db *gorm.DB, err error) {

	//db, err = gorm.Open("postgres", "user=postgres dbname=plant sslmode=disable password=postgres")
	db, err = gorm.Open(conf.Dialect, "user=postgres dbname=plant sslmode=disable password=postgres")
	if err != nil {
		logging.Log.Errorf("Error during connection to Postgres has occurred %s", err.Error())
		return nil, err
	} else {
		logging.Log.Info("Connection to Pg has been established")
	}
	return db, err
}

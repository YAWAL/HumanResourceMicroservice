package database

import (
	"fmt"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func PGconn(conf Config) (db *gorm.DB, err error) {
	db, err = gorm.Open(conf.Dialect, fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s", conf.User,
		conf.DataBaseName, conf.SSLMode, conf.Password))
	if err != nil {
		logging.Log.Errorf("error during connection to Postgres has occurred %s", err.Error())
		return nil, err
	} else {
		logging.Log.Info("connection to Postgres has been established")
	}
	return db, err
}

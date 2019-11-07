package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoDBstring = "mongodb://%v:%v"

// Mongo creates client to MongoDB with host and port provided in database config file's part
func Mongo(conf MongoConfig) (mongoCl *mongo.Client, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI(conf.Host, conf.Port)))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}
	return client, nil
}

func mongoURI(host, port string) string {
	return fmt.Sprintf(mongoDBstring, host, port)
}

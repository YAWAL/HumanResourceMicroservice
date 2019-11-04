package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo creates client to MongoDB with host and port provided in database config file part
func Mongo(conf MongoConfig) (mongoCl *mongo.Client, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI(conf.Host, conf.Port)))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		return nil, err
	}
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}
	return client, nil
}

func mongoURI(host, port string) string {
	return fmt.Sprintf("mongodb://%v:%v", host, port)
}

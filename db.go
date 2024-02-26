package main

import (
	"assignment-permission/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB(ctx context.Context) (err error) {
	config.DBConnection, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Config.MongodbConnectionString))
	if err != nil {
		return err
	}
	return nil
}

func fetchDocument(res *document, collection string) (err error) {
	return nil
}

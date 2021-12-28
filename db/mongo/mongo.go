package mongo

import (
	"context"
	"log"
	"web/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mgoCli *mongo.Client
	err    error
)

func Instance() *mongo.Client {
	config.Init()
	conf := config.GetConfig()
	clientOptions := options.Client().ApplyURI(conf.Mongo.Uri)
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return mgoCli
}

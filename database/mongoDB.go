package database

import (
	"context"
	"fmt"
	config "golang-beer-example/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientMongo *mongo.Database

func InitMongoDB() (err error) {
	if clientMongo != nil {
		clientMongo = nil
	}

	option := options.Client()
	option.SetMaxPoolSize(10)
	option.SetMinPoolSize(5)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, option.ApplyURI(config.AppConfig.DB.MongoDB.URI))

	if err != nil {
		fmt.Println("Connect MongoDB Failed")
		defer cancel()
		return
	}
	defer cancel()
	clientMongo = client.Database(config.AppConfig.DB.MongoDB.DBName)

	fmt.Println("Initial MongoDB Success")

	return
}

func GetMongoDbPool() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := clientMongo.Client().Ping(ctx, nil); err != nil {
		InitMongoDB()
	}
	return clientMongo
}

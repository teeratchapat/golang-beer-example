package main

import (
	"context"
	"fmt"
	"log"

	config "golang-beer-example/configs"
	"golang-beer-example/modules/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	config.LoadConfig()

	test := config.AppConfig.DB.Maria.Username
	fmt.Println(test)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.DB.Maria.Username,
		config.AppConfig.DB.Maria.Password,
		config.AppConfig.DB.Maria.Host,
		config.AppConfig.DB.Maria.Port,
		config.AppConfig.DB.Maria.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MariaDB: %v", err)
	}

	if err := db.AutoMigrate(&models.Beer{}); err != nil {
		log.Fatalf("MariaDB migration failed: %v", err)
	}
	log.Println("MariaDB migration completed successfully")

	mongoURI := config.AppConfig.DB.MongoDB.URI
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	mongoDB := mongoClient.Database("beer_logs")
	logCollection := mongoDB.Collection("logs")

	_, err = logCollection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{{Key: "beer_id", Value: 1}},
	})
	if err != nil {
		log.Fatalf("Failed to create MongoDB indexes: %v", err)
	}

	log.Println("MongoDB migration completed successfully")
}

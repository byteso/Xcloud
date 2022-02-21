package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/byteso/Xcloud/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database *mongo.Database
	Client   *mongo.Client
	ctx      context.Context
)

func InitEngine() {
	Database = initDatabase()
}

func initDatabase() *mongo.Database {
	c := config.Config

	mongoUrl := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", c.Data.Mongo.User, c.Data.Mongo.Password, c.Data.Mongo.Ip, c.Data.Mongo.Port, c.Data.Mongo.DatabaseName)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(c.Data.Mongo.DatabaseName)

	return database
}

func Close() {
	Client.Disconnect(ctx)
}

package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/byteso/Xcloud/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	ctx    context.Context
)

func InitEngine() {
	Client = initDatabase()
	return
}

func initDatabase() *mongo.Client {
	c := configs.Config

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

	return client
}

func Close() {
	Client.Disconnect(ctx)
}

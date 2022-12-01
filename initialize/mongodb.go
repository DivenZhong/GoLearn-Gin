package initialize

import (
	"GinDemo_v1/global"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func initMongoDb() {
	mongoDbConfig := global.GvaConfig.MongoDb
	if !mongoDbConfig.Enable {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// mongodb://root:go123#@127.0.0.1:27017/gin_demo
	mongoDbUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		mongoDbConfig.User, mongoDbConfig.Password, mongoDbConfig.Host, mongoDbConfig.Port, mongoDbConfig.Database)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUrl))
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal("connect ot mongodb failed: ", err)
	}
	log.Println("Connected to MongoDB")
	global.GvaMongoDbClient = client
}

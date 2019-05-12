package database

import (
	"context"
	"github.com/aolab/letsgo/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"strconv"
)

func New(ctx context.Context, configuration conf.Configuration) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+configuration.DatabaseAddress+":"+strconv.Itoa(configuration.DatabasePort)))
	if err != nil {
		log.Fatal(err)
	}
	isAlive := IsAlive(client, ctx)
	if !isAlive {
		log.Fatal(err)
	}
	return client
}

func GetCollection(c *mongo.Client, database string, collectionName string) *mongo.Collection {
	collection := c.Database(database).Collection(collectionName)
	if collection == nil {
		log.Fatal("Database Connection Failed")
	}
	return collection
}

func IsAlive(client *mongo.Client, ctx context.Context) bool {
	err := client.Ping(ctx, readpref.Primary())
	return err == nil
}

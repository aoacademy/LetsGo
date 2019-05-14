package main

import (
	"context"
	"github.com/aolab/letsgo/conf"
	"github.com/aolab/letsgo/database"
	"github.com/aolab/letsgo/router"
)

func main() {
	configuration := conf.New()
	ctx := context.Background()
	client := database.New(ctx, configuration)
	collection := database.GetCollection(client, configuration.Database, configuration.Collection)
	routing := router.AddPaths(collection, ctx)
	router.BindHttp(routing, configuration)
}

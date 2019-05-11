package main

import (
	"context"
	"time"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type test struct {
  lat  string `json:"name" form:"name" query:"name"`
  lng string `json:"email" form:"email" query:"email"`
}

func JsonHandler(c echo.Context) (err error) {
  t := new(test)
  if err = c.Bind(t); err == nil{
    return err
  }
  client, err := mongo.NewClient(options.Client().ApplyURI("localhost:27017"))
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  collection := client.Database("testing").Collection("lat_lng")
  res, err := collection.InsertOne(ctx,t)
  id := res.InsertedID
  return c.String(http.StatusOK, fmt.Sprintf("man amadam with id: %s", id))
}

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.POST("/posting",JsonHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

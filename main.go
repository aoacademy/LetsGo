package main

import (
	"./conf"
	"./models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tkanos/gonfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"strconv"
)

func main() {
	configuration := conf.Configuration{}
	_ = gonfig.GetConf("conf.json", &configuration)

	ctx := context.Background() ///WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(configuration.Database).Collection(configuration.Collection)
	if collection == nil {
		log.Fatal("Database Connection Failed")
	}
	httpBind(configuration, collection, ctx)
}

//TODO needs refactor
func httpBind(configuration conf.Configuration, collection *mongo.Collection, ctx context.Context) {
	e := echo.New()
	e.GET("/ping", ping)
	e.POST("/insert", func(c echo.Context) error {
		location := new(models.Location)
		if err := c.Bind(location); err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{false, "", "Bad Request!"})
		}
		insertResult, err := collection.InsertOne(ctx, bson.M{"lat": location.Latitude, "lng": location.Longitude})
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, models.Response{false, "", "Service Unavailable. Try Again!"})
		} else {
			insertedID := insertResult.InsertedID.(primitive.ObjectID).Hex()
			return c.JSON(http.StatusCreated, models.Response{true, insertedID, ""})
		}
	})
	e.GET("/getAll", func(c echo.Context) error {
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		var arr []models.Location
		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			latInter, lngInter := result["lat"], result["lng"]
			if latInter != nil && lngInter != nil {
				lat, _ := strconv.ParseFloat(fmt.Sprintf("%v", latInter), 64)
				lng, _ := strconv.ParseFloat(fmt.Sprintf("%v", lngInter), 64)
				arr = append(arr, models.Location{lat, lng})
			}
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		var jsonData []byte
		jsonData, err = json.Marshal(models.GetAllResponse{true, arr})
		if err != nil {
			log.Println(err)
		}
		return c.String(http.StatusOK, string(jsonData))
	})
	e.GET("/get/:id", getById)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(configuration.Port)))
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

//TODO Implement this
func getById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "TODO "+id)
}

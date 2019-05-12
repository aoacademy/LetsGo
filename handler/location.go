package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aolab/letsgo/messages"
	"github.com/aolab/letsgo/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

func GetAll(c echo.Context, collection *mongo.Collection, ctx context.Context) error {
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
			arr = append(arr, models.Location{Latitude: lat, Longitude: lng})
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var jsonData []byte
	jsonData, err = json.Marshal(models.Response{Ok: true, Result: arr})
	if err != nil {
		log.Println(err)
	}
	return c.String(http.StatusOK, string(jsonData))
}

func Insert(c echo.Context, collection *mongo.Collection, ctx context.Context) error {
	location := new(models.Location)
	if err := c.Bind(location); err != nil {
		return c.JSON(http.StatusBadRequest, messages.BadRequest)
	}
	insertResult, err := collection.InsertOne(ctx, bson.M{"lat": location.Latitude, "lng": location.Longitude})
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, messages.ServiceUnavailable)
	} else {
		insertedID := insertResult.InsertedID.(primitive.ObjectID).Hex()
		return c.JSON(http.StatusCreated, models.Response{Ok: true, Result: insertedID})
	}
}

//TODO implement this
func GetById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "TODO "+id)
}

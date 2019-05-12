package handler

import (
	"context"
	"encoding/json"
	"github.com/aolab/letsgo/messages"
	"github.com/aolab/letsgo/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func GetAll(c echo.Context, collection *mongo.Collection, ctx context.Context) error {
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var result []models.Location
	for cur.Next(ctx) {
		var next models.Location
		err := cur.Decode(&next)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, next)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var jsonData []byte
	jsonData, err = json.Marshal(models.Response{Ok: true, Result: result})
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
	insertResult, err := collection.InsertOne(ctx, models.Location{Latitude: location.Latitude, Longitude: location.Longitude})
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

package main

import (
	"context"
	"encoding/json"
	"github.com/aolab/letsgo/conf"
	"github.com/aolab/letsgo/database"
	"github.com/aolab/letsgo/models"
	"github.com/aolab/letsgo/router"
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIntegrateRouter(t *testing.T) {
	configuration := conf.NewTest()
	ctx := context.Background()
	client := database.New(ctx, configuration)
	collection := database.GetCollection(client, configuration.Database, configuration.Collection)
	routing := router.AddPaths(collection, ctx)
	ping(routing, t)
	insert(routing, t)
	getAll(routing, t)
}

func ping(routing *echo.Echo, t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	rr := httptest.NewRecorder()
	routing.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ping Failed.")
	}
}

func insert(routing *echo.Echo, t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var jsonData []byte
	jsonData, err := json.Marshal(models.Location{Latitude: r1.Float64(), Longitude: r1.Float64()})
	if err != nil {
		log.Println(err)
	}
	req, _ := http.NewRequest("POST", "/insert", strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	routing.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Insert Failed.")
	}
}

func getAll(routing *echo.Echo, t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomLat := r1.Float64()
	randomLong := r1.Float64()
	var jsonData []byte
	jsonData, err := json.Marshal(models.Location{Latitude: randomLat, Longitude: randomLong})
	if err != nil {
		log.Println(err)
	}
	req, _ := http.NewRequest("POST", "/insert", strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	routing.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Insert Failed.")
	}

	req, _ = http.NewRequest("GET", "/getAll", nil)
	rr = httptest.NewRecorder()
	routing.ServeHTTP(rr, req)

	decoder := json.NewDecoder(rr.Body)
	var decodedRequest models.Response
	err = decoder.Decode(&decodedRequest)
	if err != nil {
		panic(err)
	}
	locations := decodedRequest.Result.([]interface{})
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	isSet := false
	for _, v := range locations {
		loc := v.(map[string]interface{})
		if loc["lat"].(float64) == randomLat && loc["lng"].(float64) == randomLong {
			isSet = true
		}
	}
	if !isSet {
		t.Errorf("Insert And Result of GetAll is wrong.")
	}
}

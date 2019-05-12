package router

import (
	"context"
	"github.com/aolab/letsgo/conf"
	"github.com/aolab/letsgo/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

func AddPaths(collection *mongo.Collection, ctx context.Context) *echo.Echo {
	e := New()
	e.GET("/ping", handler.Ping)
	e.POST("/insert", func(c echo.Context) error {
		return handler.Insert(c, collection, ctx)
	})
	e.GET("/getAll", func(c echo.Context) error {
		return handler.GetAll(c, collection, ctx)
	})
	e.GET("/get/:id", handler.GetById)
	return e
}

func BindHttp(e *echo.Echo, configuration conf.Configuration) {
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(configuration.Port)))

}

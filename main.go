package main

import (
	"github.com/aolab/letsgo/conf"
	"github.com/labstack/echo/v4"
	"github.com/tkanos/gonfig"
	"net/http"
	"strconv"
)

func main() {
	configuration := conf.Configuration{}
	gonfig.GetConf("conf.json", &configuration)

	e := echo.New()
	e.GET("/ping", ping)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(configuration.Port)))
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

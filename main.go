package main

import (
        "net/http"

        "github.com/labstack/echo"
)

func main() {
        e := echo.New()
        e.GET("/ping", func(c echo.Context) error {
                return c.String(http.StatusOK, "")
        })
        e.Logger.Fatal(e.Start(":1322"))
}

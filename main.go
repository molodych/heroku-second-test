package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Привет тебе от Игоря!")
	})
	port := ":" + os.Getenv("PORT")
	fmt.Println(port)
	e.Logger.Fatal(e.Start(port))
}

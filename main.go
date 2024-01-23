package main

import (
	"go-shorturl/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handlers.ShortUrl(e)
	e.Static("/", "views")
	e.File("/", "views/index.html")
	e.Logger.Fatal(e.Start(":9090"))
}
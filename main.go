package main

import (
	"go-shorturl/handlers"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("error loading file .env")
	}

	handlers.ShortUrl(e)
	e.Static("/", "views")
	e.File("/index", "views/index.html")
	e.Logger.Fatal(e.Start(":9090"))
}

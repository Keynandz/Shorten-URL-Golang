package main

import (
	"fmt"
	"go-shorturl/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("error loading file .env")
	}

	format := ":%s"
	port := fmt.Sprintf(format, os.Getenv("PORT"))
	handlers.ShortUrl(e)
	e.Static("/", "views")
	e.File("/index", "views/index.html")
	e.Logger.Fatal(e.Start(port))
}

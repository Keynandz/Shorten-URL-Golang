package main

import (
	"fmt"
	"go-shorturl/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("error loading file .env")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Auth-Token"},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	format := ":%s"
	port := fmt.Sprintf(format, os.Getenv("PORT"))
	handlers.ShortUrl(e)
	e.Static("/", "views")
	e.File("/index", "views/index.html")
	e.Logger.Fatal(e.Start(port))
}

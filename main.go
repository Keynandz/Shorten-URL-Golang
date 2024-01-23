package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-shorturl/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.GET("/", handlers.Home)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	format := "0.0.0.0:%s"
	port := fmt.Sprintf(format, os.Getenv("PORT"))
	e.Start(port)
}

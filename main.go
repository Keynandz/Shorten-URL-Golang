package main

import (
	"fmt"
	"go-shorturl/repositories"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var urlMap = make(map[string]string)

type ShortenRequest struct {
	LongURL         string `json:"longURL"`
	CustomShortURL  string `json:"customShortURL"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortURL"`
}

func main() {
	e := echo.New()
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("error loading file .env")
	}

	e.POST("keynandz/shorten", func(c echo.Context) error {
		req := new(ShortenRequest)
		if err := c.Bind(req); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request body")
		}

		var shortURL string
		if req.CustomShortURL != "" {
			shortURL = req.CustomShortURL
		} else {
			shortURL = repositories.GenerateShortURL()
		}

		urlMap[shortURL] = req.LongURL

		format := "http://%s:%s/keynandz/%s"
		response := fmt.Sprintf(format, os.Getenv("HOST"), os.Getenv("PORT"), shortURL)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"short link":  response,
			"status_code": http.StatusOK,
		})
	})

	e.GET("keynandz/:shortURL", func(c echo.Context) error {
		shortURL := c.Param("shortURL")
		longURL, exists := urlMap[shortURL]
		if exists {
			return c.Redirect(http.StatusMovedPermanently, longURL)
		}
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Short URL not found",
			"error":   http.StatusNotFound,
		})
	})

	e.Static("/", "views")
	e.File("/index", "views/index.html")

	e.Logger.Fatal(e.Start(":9090"))
}

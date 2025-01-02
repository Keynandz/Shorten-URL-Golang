package handlers

import (
	"fmt"
	"go-shorturl/repositories"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

var urlMap = make(map[string]string)

type ShortenRequest struct {
	LongURL        string `json:"longURL"`
	CustomShortURL string `json:"customShortURL"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortURL"`
}

func ShortUrl(e *echo.Echo) {
	e.POST("/keynandz/shorten", func(c echo.Context) error {
		req := new(ShortenRequest)
		if err := c.Bind(req); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request body")
		}

		// Validasi spasi dalam CustomShortURL
		if strings.Contains(req.CustomShortURL, " ") {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Custom short URL cannot contain spaces",
			})
		}

		var shortURL string
		if req.CustomShortURL != "" {
			// Cek jika CustomShortURL sudah ada
			if _, exists := urlMap[req.CustomShortURL]; exists {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "Custom short URL already exists, please choose another one",
				})
			}
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

	e.GET("/keynandz/:shortURL", func(c echo.Context) error {
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
}

func Home(c echo.Context) error {
	content, err := os.ReadFile("index.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Could not open requested file")
	}

	return c.HTML(http.StatusOK, string(content))
}

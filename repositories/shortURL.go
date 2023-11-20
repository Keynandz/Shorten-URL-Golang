package repositories

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateShortURL() string {
	const length = 6
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	var result strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		result.WriteString(string(charset[randomIndex]))
	}

	return result.String()
}

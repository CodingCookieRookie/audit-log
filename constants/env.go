package constants

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func GetAppSecret() string {
	return os.Getenv("APP_SECRET")
}

func GetAPIKey() string {
	return os.Getenv("API_KEY")
}

func GetServiceEmailPassword() string {
	return os.Getenv("EMAIL_PASSWORD")
}

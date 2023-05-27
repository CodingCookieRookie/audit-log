package constants

import (
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()
}

func GetAPIKey() string {
	return os.Getenv("API_KEY")
}

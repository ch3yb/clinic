package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type databaseConfig struct {
	URL  string
	FILE string
}

type config struct {
	HttpPort string
	Database *databaseConfig
}

var Conf *config

func init() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
		return
	}

	Conf = &config{
		HttpPort: os.Getenv("HTTP_PORT"),
		Database: &databaseConfig{
			os.Getenv("DATABASE_URL"),
			os.Getenv("DATABASE_MIGRATION_FILE"),
		},
	}
}

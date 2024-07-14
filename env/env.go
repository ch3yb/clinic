package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SslMode  string
}

type Config struct {
	HttpPort string
	Database *DatabaseConfig
}

var Conf *Config

func init() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
		return
	}

	Conf = &Config{
		HttpPort: os.Getenv("HTTP_PORT"),
		Database: &DatabaseConfig{
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USERNAME"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DBNAME"),
			os.Getenv("POSTGRES_SSL_MODE"),
		},
	}
}

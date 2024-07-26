package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var (
	PORT         string
	DATABASE_URI string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3001
	}

	PORT = fmt.Sprintf(":%d", port)
	DATABASE_URI = os.Getenv("DATABASE_URI")
}

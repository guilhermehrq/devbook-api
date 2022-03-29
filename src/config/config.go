package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	APIPort          = 0
)

func LoadEnv() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println("Error to convert var 'API_PORT', setting default port to 5000")
		APIPort = 5000
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
}

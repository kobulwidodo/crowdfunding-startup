package main

import (
	"bwastartup/config"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var APP_ENV = os.Getenv("APP_ENV")
	if APP_ENV != "PRODUCTION" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Setup Database
	_, err := config.InitDB()
	if err != nil {
		fmt.Println(err)
		panic("Something wrong with database")
	}
	fmt.Println("Sukses connect ke database!")
}

package main

import (
	"bwastartup/config"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Setup Environment
	var APP_ENV = os.Getenv("APP_ENV")
	if APP_ENV != "PRODUCTION" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Setup Database
	db, err := config.InitDB()
	if err != nil {
		fmt.Println(err)
		panic("Something wrong with database")
	}
	fmt.Println("Sukses connect ke database!")

	// Setup Repository
	userRepository := user.NewRepository(db)

	// Setup Service
	userSevice := user.NewService(userRepository)

	// Setup Handler
	userHandler := handler.NewUserHandler(userSevice)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users", userHandler.RegisterUser)
	}

	r.Run()
}

package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/config"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// Setup Service
	userSevice := user.NewService(userRepository)
	authService := auth.NewJwtService()
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository)

	// Setup Handler
	userHandler := handler.NewUserHandler(userSevice, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := gin.Default()

	r.Static("/images", "./images")
	api := r.Group("/api")
	{
		api.POST("/users", userHandler.RegisterUser)
		api.POST("/sessions", userHandler.Login)
		api.POST("/email_checker", userHandler.CheckEmailAvailability)
		api.POST("/avatars", authMiddleware(authService, userSevice), userHandler.UploadAvatar)

		api.GET("/campaigns", campaignHandler.GetCampaigns)
		api.GET("/campaigns/:id", campaignHandler.GetCampaign)
		api.POST("/campaigns", authMiddleware(authService, userSevice), campaignHandler.CreateCampaign)
		api.PUT("/campaigns/:id", authMiddleware(authService, userSevice), campaignHandler.UpdateCampaign)
		api.POST("/campaign-images", authMiddleware(authService, userSevice), campaignHandler.UploadImage)

		api.GET("/campaigns/:id/transactions", authMiddleware(authService, userSevice), transactionHandler.GetTransactionByCampaignId)
	}

	r.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Not a Bearer token", http.StatusUnauthorized, "gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Token Invalid", http.StatusUnauthorized, "gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		var user user.User
		user, err = userService.GetUserById(userId)
		if err != nil {
			response := helper.ApiResponse("Failed to get user", http.StatusUnauthorized, "gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("userLoggedin", user)
	}
}

package main

import (
	"AliveVirtualGift_RestAPI/src/handler"
	"AliveVirtualGift_RestAPI/src/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	// Load environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	router := gin.Default()

	router.POST("/login", handler.Login)
	router.POST("account/create", handler.Create)

	authorized := router.Group("/")
	authorized.Use(middleware.TokenAuth)
	{
		router.POST("/logout", handler.Logout)
		router.GET("/account/show", handler.Show)
		router.PUT("/account/update", handler.Update)
	}

	log.Fatal(router.Run(":" + os.Getenv("REST_API_PORT")))
}

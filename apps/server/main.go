package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/scallyt/Edina-Neni-School/services"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//	ROUTES
	router.GET("/tasks", services.TaskList)

	router.POST("/user/login", services.UserLogin)

	router.Run(os.Getenv("PORT"))
}

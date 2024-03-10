package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/scallyt/Edina-Neni-School/middlewares"
	"github.com/scallyt/Edina-Neni-School/services"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	router := gin.Default()

	router.Use(middlewares.TokenAuthMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//	ROUTES
	router.GET("/tasks", services.TaskList)
	router.POST("/tasks", services.TaskCreate)
	router.PATCH("/tasks/:id", services.TaskUpdate)
	router.DELETE("/tasks/:id", services.TaskDelete)

	router.POST("/user/login", services.UserLogin)
	router.GET("/users", services.UserLists)

	router.Run(":4040")
}

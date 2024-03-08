package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scallyt/Edina-Neni-School/services"
)


func main() {

	

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//	ROUTES
	router.POST("/user/login", services.UserLogin)

	router.Run(":4040")
}

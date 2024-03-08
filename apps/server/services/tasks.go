package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tasks struct {
	Id   uint   `json:"id"`
	Date string `json:"date"`
	Text string `json:"text"`
}

func TaskList(c *gin.Context) {
	var tasks []Tasks
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

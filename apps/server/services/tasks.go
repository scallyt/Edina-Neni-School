package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scallyt/Edina-Neni-School/models"
)

func TaskList(c *gin.Context) {
	var task []models.Tasks
	db.Find(&task)

	c.JSON(http.StatusOK, gin.H{"tasks": task})
}

func TaskCreate(c *gin.Context) {
	var task models.Tasks
	c.BindJSON(&task)
	db.Create(&task)

	c.JSON(http.StatusCreated, gin.H{"task": task})
}

func TaskUpdate(c *gin.Context) {
	var task models.Tasks

	id := c.Param("id")
	db.First(&task, id)

	c.BindJSON(&task)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func TaskDelete(c *gin.Context) {
	var task models.Tasks

	id := c.Param("id")
	db.First(&task, id)

	db.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"task": id})
}

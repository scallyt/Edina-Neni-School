package services

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scallyt/Edina-Neni-School/internal"
	"github.com/scallyt/Edina-Neni-School/models"
)

var user internal.User

var ResponseUser = models.responsUser{
	int(user.ID),
	user.Email,
	user.Token,
}

func UserLogin(c *gin.Context) {
	var data models.LoginData
	var err error
	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash := md5.Sum([]byte(data.Password))
	hashedUserPassword := hex.EncodeToString(hash[:])

	if err := db.Where("email = ? AND password = ? AND username = ?", data.Email, hashedUserPassword, data.UserName).First(&user).Error; err != nil {
		// User not found or error occurred
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password or email"})
		return
	}

	// Login successful
	c.JSON(http.StatusOK, gin.H{"user": ResponseUser})
}

func UserLists(c *gin.Context) {
	var users []internal.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": ResponseUser})
}

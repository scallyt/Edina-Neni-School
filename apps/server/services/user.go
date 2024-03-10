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

	responseUser := models.ResponseUser{
		ID:    int(user.ID),
		Email: user.Email,
		Token: user.Token,
	}

	// Login successful
	c.JSON(http.StatusOK, gin.H{"user": responseUser})
}

func UserLists(c *gin.Context) {
	var users []internal.User
	db.Find(&users)

	responseUser := make([]models.ResponseUser, len(users))
	for i, u := range users {
		responseUser[i] = models.ResponseUser{
			ID:    int(u.ID),
			Email: u.Email,
			Token: u.Token,
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": responseUser})
}

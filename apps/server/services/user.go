package services

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scallyt/Edina-Neni-School/internal"
	"gorm.io/gorm"
)

var db *gorm.DB

type LoginData struct {
	UserName string `json:"username"` // Capitalized field name
	Email    string `json:"email"`    // Capitalized field name
	Password string `json:"password"` // Capitalized field name
}

type responsUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func init() {
	db = internal.ConnectDatabase()
}

func UserLogin(c *gin.Context) {
	var data LoginData
	var err error
	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash := md5.Sum([]byte(data.Password))
	hashedUserPassword := hex.EncodeToString(hash[:])

	var user internal.User
	if err := db.Where("email = ? AND password = ?", data.Email, hashedUserPassword).First(&user).Error; err != nil {
		// User not found or error occurred
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid asd"})
		return
	}

	var ResponseUser = responsUser{
		int(user.ID),
		user.Email,
		user.Token,
	}

	// Login successful
	c.JSON(http.StatusOK, gin.H{"user": ResponseUser})
}

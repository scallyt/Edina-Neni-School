package internal

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Token    string `gorm:"unique;not null`

	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"`
}

type Tasks struct {
	ID   uint   `gorm:"primaryKey"`
	Date string `gorm:"unique;not null`
	Text string `gorm:"unique;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"`
}

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, err := gorm.Open(sqlite.Open("./dev.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{}, &Tasks{})
	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {

		defaultUser := &User{
			Email:    os.Getenv("Email"),
			Username: os.Getenv("Username"),
			Password: os.Getenv("Password"),
			Token:    os.Getenv("Token"),
		}
		db.Create(defaultUser)
	}

	return db
}

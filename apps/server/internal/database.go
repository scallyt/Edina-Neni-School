package internal

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey;default:1"`
	Email    string `gorm:"unique;not null;default:test@gmail.com"`
	Username string `gorm:"unique;not null;default:edina_neni"`
	Password string `gorm:"not null;default:123"`
	Token    string `gorm:"unique;not null;default:bz1Ye6z5tXwFu9ZUe8kqjW3p"`

	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"`
}

type Tasks struct {
	ID   uint   `gorm:"primaryKey;default:1"`
	Date string `gorm:"unique;not null`
	Text string `gorm:"unique;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"`
}

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./dev.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{}, &Tasks{})
	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {

		defaultUser := &User{
			Email:    "test@gmail.com",
			Username: "edina_neni",
			Password: "0b95527e8d68b6afaee0de6d574be951",
			Token:    "bz1Ye6z5tXwFu9ZUe8kqjW3p",
		}
		db.Create(defaultUser)
	}

	return db
}

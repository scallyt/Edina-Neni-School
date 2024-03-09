package services

import (
	"github.com/scallyt/Edina-Neni-School/internal"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = internal.ConnectDatabase()
}

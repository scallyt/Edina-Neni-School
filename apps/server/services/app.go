package services

import (
	"github.com/scallyt/Edina-Neni-School/internal"
	"gorm.io/gorm"
)

var db *gorm.DB
var user internal.User

func init() {
	db = internal.ConnectDatabase()
}

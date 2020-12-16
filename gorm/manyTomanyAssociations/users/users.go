package users

import (
	"github.com/jinzhu/gorm"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}
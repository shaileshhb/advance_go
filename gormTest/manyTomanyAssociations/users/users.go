package users

import (
	"gormTest/manyTomanyAssociations/languages"

	"github.com/jinzhu/gorm"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	gorm.Model
	Name      string
	Languages []languages.Language `gorm:"many2many:user_languages;"`
}

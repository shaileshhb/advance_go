package users

import (
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/advance_go/gormTest/manyTomanyAssociations/languages"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	gorm.Model
	Name      string
	Languages []languages.Language `gorm:"many2many:user_languages;"`
}

package languages

import (
	"gormTest/manyTomanyAssociations/users"

	"github.com/jinzhu/gorm"
)

type Language struct {
	gorm.Model
	Name  string
	Users []*users.User `gorm:"many2many:user_languages;"`
}

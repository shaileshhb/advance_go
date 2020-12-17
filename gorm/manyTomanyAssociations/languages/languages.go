package languages

import (
	"github.com/jinzhu/gorm"
)

type Language struct {
	gorm.Model
	Name string
}

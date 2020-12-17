package customer

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name    string
	Address string
	Age     int
}

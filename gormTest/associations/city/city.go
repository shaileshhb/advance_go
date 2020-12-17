package city

import (
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/advance_go/gormTest/associations/location"
)

type City struct {
	gorm.Model
	CityName  string
	Locations []location.Location `gorm:"ForeignKey:CityID"`
}

package city

import (
	"gormTest/associations/location"

	"github.com/jinzhu/gorm"
)

type City struct {
	gorm.Model
	CityName  string
	Locations []location.Location `gorm:"ForeignKey:CityID"`
}

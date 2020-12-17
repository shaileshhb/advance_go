package associationsCrud

import (
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/advance_go/gormTest/associations/city"
	"github.com/shaileshhb/advance_go/gormTest/associations/location"
)

func CountAssociations(db *gorm.DB, id uint) int {
	return db.Debug().First(&city.City{}, id).Association("Locations").Count()
}

func FindAssociatons(db *gorm.DB, id uint) []location.Location {
	locationsFound := []location.Location{}

	db.Debug().First(&city.City{}, id).Association("Locations").Find(&locationsFound)
	return locationsFound
}

func ReplaceAssociations(db *gorm.DB, id uint) {

	replaceLocations := []location.Location{
		{
			LocationName: "Gateway Of India",
		},
		{
			LocationName: "Location2",
		},
	}

	db.Debug().First(&city.City{}, id).Association("Locations").Replace(replaceLocations)

}

func RemoveAssociations(db *gorm.DB, id uint) {

	findLocations := FindAssociatons(db, id)

	db.Debug().First(&city.City{}, id).Association("Locations").Delete(findLocations)
}

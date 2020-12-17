package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/advance_go/gormTest/associations/associationsCrud"
	"github.com/shaileshhb/advance_go/gormTest/associations/city"
	"github.com/shaileshhb/advance_go/gormTest/associations/location"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&city.City{}, &location.Location{})
	// db.Model(&location.Location{}).AddForeignKey("city_id", "locations(id)", "CASCADE", "RESTRICT")

	// var pun = &city.City{
	// 	CityName: "Amritsar",
	// 	Locations: []location.Location{
	// 		{
	// 			LocationName: "Golden Temple",
	// 		},
	// 		{
	// 			LocationName: "Jallianwala Bagh",
	// 		},
	// 	},
	// }

	// db.Create(pun)
	// fmt.Println("Data added successfully")

	// Count Associations for locations with city_id = 2
	// fmt.Println("Associations Count:", associationsCrud.CountAssociations(db, 2))

	// Find Associations for location with city_id = 3
	// fmt.Println("Slice of Associations:", associationsCrud.FindAssociatons(db, 3))

	// Remove Associations for location with city_id = 3
	// associationsCrud.RemoveAssociations(db, 3)
	// fmt.Println("Associations Removed")

	// Replace Associations for location with city_id = 1
	associationsCrud.ReplaceAssociations(db, 1)
	fmt.Println("Associations Replaced")
}

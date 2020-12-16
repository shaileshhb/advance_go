package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Users struct {
	ID      int    `gorm:"primary_key"; "AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open", err)
		return
	}

	user := &Users{ID: 4}

	// works on primary key
	// select, edit, and save
	// db.Find(&user)
	// user.Address = "Goa"
	// db.Save(&user)

	// update with column names, and not attribute
	db.Model(&user).Update("Address", "Delhi")

	db.Model(&user).Updates(
		map[string]interface{}{
			"Name":    "Ben Stokes",
			"Address": "Delhi",
		})

	// UpdateColumn()
	// db.Model(&user).UpdateColumn("Address", "Chandigarh")

	// UpdateColumns()
	// db.Model(&user).UpdateColumns(
	// 	map[string]interface{}{
	// 		"Name":    "Tom Shacks",
	// 		"Address": "Chandigarh",
	// 	})

	// using Find()
	// db.Find(&user).Update("Address", "Manali")

	// Batch update not working
	// db.Table("gorm_test").Where("ID = ?", 4).Update("Address", "Chennai")

	fmt.Println(user, &user)
	fmt.Println("User successfully updated")

}

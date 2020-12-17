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

	// Connect to db
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open", err)
		return
	}

	// Selects record and deletes it
	// db.Table("users").Where("id=?", 3).Delete(&Users{})

	// Find the record and delete it
	// db.Where("Address=?", "Pune").Delete(&Users{})

	//Select all records from a model and delete all
	// db.Model(&Users{}).Delete(&Users{})

	fmt.Println("User Deleted")
}

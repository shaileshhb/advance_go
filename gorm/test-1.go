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
	// creates new table
	// db.CreateTable(&Users{})

	// insert data to table
	// INSERT INTO Users (name, address) VALUES ('Ben', 'Mumbai')
	// user1 := &Users{
	// 	Name:    "Ben",
	// 	Address: "Mumbai",
	// }

	// db.Create(user1)

	var user []Users = []Users{
		Users{Name: "Tom", Address: "Pune"},
		Users{Name: "Sam", Address: "Banglore"},
		Users{Name: "Adam", Address: "Goa"},
	}

	for _, u := range user {
		db.Create(&u)
		fmt.Println("data successfully inserted")
	}

	fmt.Println("All data successfully inserted")

}

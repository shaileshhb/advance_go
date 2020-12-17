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
	db.CreateTable(&Users{})

	fmt.Println("Table successfully created")

}

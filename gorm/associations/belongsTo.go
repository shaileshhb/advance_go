package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// User belongs to `Company`, `CompanyID` is the foreign key
type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open", err)
		return
	}
	fmt.Println("Db successfully connected")

	user1 := User{}
	company1 := Company{}

	db.AutoMigrate(&user1, &company1)

	fmt.Println(queryForCompany(db, 101))
}

func queryForCompany(db *gorm.DB, id int) Company {
	company := Company{}

	db.First(&User{}, id).Related(&company)

	return company
}

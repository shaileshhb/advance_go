package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open", err)
		return
	}

	fmt.Println("First Row:", firstQueryFromDB(db))
	fmt.Println("Second Row:", lastQueryFromDB(db))
	// fmt.Println("All records:", queryAllRecords(db))
	fmt.Println("User with id=4", getSpecifiedQuery(db, 4))
}

// Get first record, order by primary key
func firstQueryFromDB(db *gorm.DB) Users {
	user := Users{}

	// SELECT * FROM users ORDER BY id LIMIT 1;
	db.First(&user)
	return user
}

// Get last record, order by primary key
func lastQueryFromDB(db *gorm.DB) Users {
	user := Users{}

	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user)
	return user
}

// Get all records
func queryAllRecords(db *gorm.DB) Users {
	user := Users{}

	// SELECT * FROM users;
	db.Find(&user)
	return user
}

// Get record with primary key (only works for integer primary key)
func getSpecifiedQuery(db *gorm.DB, id int) Users {
	user := Users{}

	// SELECT * FROM users WHERE id = id(from arguments);
	db.First(&user, id)
	return user
}

type Users struct {
	ID      int    `gorm:"primary_key"; "AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

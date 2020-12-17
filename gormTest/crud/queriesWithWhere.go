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

	fmt.Println("first matched query:", firstMatchedQuery(db, 9))
	fmt.Println("matched query:", getAllMatchedRecords(db, "Adam"))
	fmt.Println("IN:", queryIN(db))
	fmt.Println("LIKE:", queryLike(db))

}

// Get first matched record
func firstMatchedQuery(db *gorm.DB, userid int) Users {
	user := Users{}

	// SELECT * FROM users WHERE id = userid ORDER BY id LIMIT 1;
	db.Where("id=?", userid).First(&user)
	return user
}

// Get all matched records
func getAllMatchedRecords(db *gorm.DB, userName string) Users {
	user := Users{}

	// SELECT * FROM users WHERE name = 'jinzhu';
	db.Where("Name = ?", userName).Find(&user)
	return user
}

// IN
func queryIN(db *gorm.DB) Users {
	user := Users{}

	// SELECT * FROM users WHERE name <> 'Tom';
	db.Where("Name IN (?)", "Hardik").Find(&user)
	return user
}

// LIKE
func queryLike(db *gorm.DB) Users {
	user := Users{}

	// SELECT * FROM users WHERE name LIKE '%jin%';
	db.Where("name LIKE ?", "tom%").Find(&user)
	return user
}

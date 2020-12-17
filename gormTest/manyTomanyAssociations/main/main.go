package main

import (
	"fmt"
	"gormTest/manyTomanyAssociations/languages"
	"gormTest/manyTomanyAssociations/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&users.User{}, &languages.Language{})

	// Creating users
	var newUser1 = &users.User{
		Name: "User3",
		Languages: []languages.Language{
			{
				Name: "Lang2",
			},
		},
	}

	// var newUser2 = &users.User{
	// 	Name: "User2",
	// 	Languages: []*languages.Language{
	// 		{
	// 			Name: "Lang3",
	// 		},
	// 		{
	// 			Name: "Lang1",
	// 		},
	// 	},
	// }

	db.Create(newUser1)
	// db.Create(newUser2)

	// crud.AddUser(db)
	// crud.DeleteUser(db)

	fmt.Println("Data entered successfully")
}

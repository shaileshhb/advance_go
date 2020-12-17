package crud

import (
	"fmt"
	"gormTest/manyTomanyAssociations/languages"
	"gormTest/manyTomanyAssociations/users"

	"github.com/jinzhu/gorm"
)

func AddUser(db *gorm.DB) {

	var lang5 = &languages.Language{
		Name: "Lang5",
	}
	lang5.ID = 15

	user1 := &users.User{
		Name: "New User",
		Languages: []*languages.Language{
			lang5,
		},
	}

	db.Create(user1)
	fmt.Println("New User added")
}

func UpdateLang(db *gorm.DB) {

	lang := &languages.Language{}
	lang.ID = 13
	db.Debug().Model(&lang).Update("Name", "New Lang 13")

	fmt.Println("Lang Updated")
}

func DeleteUser(db *gorm.DB) {

	delete := languages.Language{}
	delete.ID = 14

	db.Debug().Model(&delete).Delete(&delete)

	fmt.Println("Lang Deleted")
}

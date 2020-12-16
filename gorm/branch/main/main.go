package main

import (
	"errors"
	"fmt"
	"gormTest/branch/associations"
	"gormTest/branch/bank"
	"gormTest/branch/customer"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open", err)
		return
	}
	defer db.Close()
	fmt.Println("Database connected successfully")

	db.AutoMigrate(&bank.Bank{}, &customer.BankCustomer{})

	// db.Model(&customer.BankCustomer{}).AddForeignKey("bank_id", "banks(id)", "CASCADE", "RESTRICT")

	// bank1 := &bank.Bank{
	// 	BankName: "DEF Bank",
	// 	Customers: []customer.BankCustomer{
	// 		{
	// 			Name:   "Sam",
	// 			Amount: 1000,
	// 		},
	// 		{
	// 			Name:   "Tom",
	// 			Amount: 5000,
	// 		},
	// 	},
	// }
	// // bank1.ID = uuid.NewV4()

	// db.Create(&bank1)
	// fmt.Println("Data successfully added")

	// cust1 := &customer.BankCustomer{}
	// db.Model(cust1).First(cust1, 501)

	// cust2 := &customer.BankCustomer{}
	// db.Model(cust1).First(cust2, 502)

	// err = SendMoney(db, cust1, cust2, 1000)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Successful...")

	// fmt.Println("Associations")
	// bank1 := bank.Bank{}
	// bankID := bank1.ID
	associations.RemoveAssociation(db)

	// associations.FindAssociations(db)
	// GetCustFromBank(db)

}

// db.Begin()
// db.Complete()
// db.Commit()

func SendMoney(db *gorm.DB, cust1 *customer.BankCustomer, cust2 *customer.BankCustomer, amt int) error {

	tx := db.Begin()

	amount := cust2.Amount - amt
	if amount <= 0 {
		tx.Rollback()
		return errors.New("Amount is greate than specified")
	}

	if err := tx.Model(cust2).Update("Amount", amount).Error; err != nil {
		tx.Rollback()

	}

	amount = cust1.Amount + amt
	if err := tx.Model(cust1).Update("Amount", amount).Error; err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

func GetCustomer(db *gorm.DB) {

	uuidID, err := uuid.FromString("24701898-d1f8-4d9a-a01d-5e2e182bf161")
	if err != nil {
		log.Fatal(err)
	}

	cust1 := customer.BankCustomer{}
	cust1.ID = uuidID

	db.Model(&cust1).First(&cust1)

	fmt.Println(cust1.BankID)

}

func GetCustFromBank(db *gorm.DB) {

	uuidID, err := uuid.FromString("20d40683-9b81-42be-b05f-bd1d31ac1743")
	if err != nil {
		log.Fatal(err)
	}

	bank1 := bank.Bank{}
	bank1.ID = uuidID
	db.Model(&bank1).Find(&bank1)

	fmt.Println(bank1)

}

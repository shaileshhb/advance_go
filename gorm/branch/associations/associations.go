package associations

import (
	"fmt"
	"gormTest/branch/bank"
	"gormTest/branch/customer"

	"github.com/jinzhu/gorm"
)

func FindAssociations(db *gorm.DB) {

	bank1 := []customer.BankCustomer{}
	// bank1.ID = bankID
	db.First(&bank.Bank{}).Association("Customers").Find(&bank1)
	fmt.Println("Bank1: ", bank1)
}

func ReplaceAssociation(db *gorm.DB) {
	
	bank1 := []customer.BankCustomer{}
	// bank1.ID =
	db.First(&bank.Bank{}).Association("Customers").Find(&bank1)

	cust1 := customer.BankCustomer{
		ID:     505,
		Name:   "Ben",
		Amount: 2000,
	}

	// cust2 := customer.BankCustomer{
	// 	ID:     504,
	// 	Name:   "Jack",
	// 	Amount: 3000,
	// }

	db.Debug().First(&bank.Bank{}).Association("Customers").Replace([]customer.BankCustomer{cust1})

	fmt.Println("Associations replaced")
}

func RemoveAssociation(db *gorm.DB) {
	
	// bank1 := []customer.BankCustomer{}
	// bank1.ID =
	// db.First(&bank.Bank{}).Association("Customers").Find(&bank1)

	cust1 := customer.BankCustomer{
		ID:     506,
		Name:   "Josh",
		Amount: 2000,
	}

	// cust2 := customer.BankCustomer{
	// 	ID:     504,
	// 	Name:   "Jack",
	// 	Amount: 3000,
	// }

	db.Debug().First(&bank.Bank{}).Association("Customers").Replace([]customer.BankCustomer{cust1})

	fmt.Println("Associations removed")
}
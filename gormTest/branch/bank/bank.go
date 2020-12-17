package bank

import (
	"gormTest/branch/customer"
	"gormTest/branch/model"
)

type Bank struct {
	model.Model
	BankName  string
	Customers []customer.BankCustomer `gorm:"ForeignKey:BankID"`
}

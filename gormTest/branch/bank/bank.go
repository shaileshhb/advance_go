package bank

import (
	"github.com/shaileshhb/advance_go/gormTest/branch/customer"
	"github.com/shaileshhb/advance_go/gormTest/branch/model"
)

type Bank struct {
	model.Model
	BankName  string
	Customers []customer.BankCustomer `gorm:"ForeignKey:BankID"`
}

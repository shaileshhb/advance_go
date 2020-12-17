package customer

import (
	"github.com/shaileshhb/advance_go/gormTest/office/model"
	"github.com/shaileshhb/advance_go/gormTest/office/order"
)

type Customer struct {
	model.Model
	Name   *string
	Orders []order.Order `gorm:"ForeignKey:CustomerID"`
}

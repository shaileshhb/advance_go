package customer

import (
	"gormTest/office/model"
	"gormTest/office/order"
)

type Customer struct {
	model.Model
	Name   *string
	Orders []order.Order `gorm:"ForeignKey:CustomerID"`
}

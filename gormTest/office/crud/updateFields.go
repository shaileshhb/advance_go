package crud

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/shaileshhb/advance_go/gormTest/office/customer"
	"github.com/shaileshhb/advance_go/gormTest/office/order"
)

func UpdateCustomer(db *gorm.DB, newName *string, id uuid.UUID) customer.Customer {

	cust := customer.Customer{}
	cust.ID = id

	customerUpdate := customer.Customer{
		Name: newName,
	}

	db.Debug().Model(&cust).Update(customerUpdate)
	return customerUpdate
}

func UpdateOrder(db *gorm.DB, newOrderName *string, id int) order.Order {

	order1 := order.Order{}
	// order1.OrderID = id

	updatedOrder := order.Order{
		OrderName: newOrderName,
	}

	db.Debug().Model(&order1).Where("order_id=?", id).Update(&updatedOrder)
	return updatedOrder
}

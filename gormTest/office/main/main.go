package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/advance_go/gormTest/office/customer"
	"github.com/shaileshhb/advance_go/gormTest/office/order"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open", err)
		return
	}
	defer db.Close()
	fmt.Println("Database connected successfully")

	db.AutoMigrate(&customer.Customer{}, &order.Order{})

	// db.Model(&order.Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "RESTRICT")

	name := "Tom"
	orderName1 := "Order3"
	orderName2 := "Order4"

	customer1 := customer.Customer{

		Name: &name,
		Orders: []order.Order{
			{
				OrderID:   503,
				OrderName: &orderName1,
			},
			{
				OrderID:   504,
				OrderName: &orderName2,
			},
		},
	}
	// customer1.ID = uuid.NewV4()
	db.Create(&customer1)
	// fmt.Println("Data successfully added")

	// Select
	// fmt.Println("Data loaded:", queryAllRecords(db))

	// Update name using pointer
	// newName := "Tom"
	// fmt.Println("Record updated:", crud.UpdateCustomer(db, &newName, 101))

	// newOrderName := "Order2"
	// fmt.Println("Order Update:", crud.UpdateOrder(db, &newOrderName, 502))

	// newOrderName = "Order3"
	// fmt.Println("Order Update:", crud.UpdateOrder(db, &newOrderName, 503))

	// newOrderName := "Order4"
	// fmt.Println("Order Update:", crud.UpdateOrder(db, &newOrderName, 504))
	// fmt.Println("Data loaded:", queryAllRecords(db))

	// var count int
	// db.Debug().Model(&customer.Customer{}).Count(&count)
	// fmt.Println("Count:", count)

}

func queryAllRecords(db *gorm.DB) customer.Customer {
	customer1 := &customer.Customer{}

	// SELECT * FROM users;
	// db.Preload("Orders").Find(&user)

	db.Preload("Orders").Find(&customer1)
	// db.Preload("Orders").Where("name=?", "Sam").First(&customer1)

	return *customer1
}

func deleteOrderRow(db *gorm.DB, id int) {

	order1 := order.Order{}
	// order1.OrderID = id

	db.Debug().Where("order_id=?", id).Delete(&order1)

	fmt.Println(order1)

}

func deleteColumn(db *gorm.DB, id uint) {

	cust := customer.Customer{}
	// cust.ID = id

	db.Debug().Where("id=?", id).Delete(&cust)
	fmt.Println(cust)
}

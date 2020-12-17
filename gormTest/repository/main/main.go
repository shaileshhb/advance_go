package main

import (
	"fmt"
	"gormTest/repository/crud"
	"gormTest/repository/customer"
	"gormTest/repository/unitOfWork"

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
	fmt.Println("DB connected Successfully")

	db.AutoMigrate(&customer.Customer{})

	uow := &unitOfWork.UnitOfWork{
		DB:        db,
		Committed: false,
		Readonly:  true,
	}

	// cust := customer.Customer{
	// 	Name:    "Jack",
	// 	Address: "delhi",
	// 	Age:     21,
	// }
	// cust.ID = 103

	// err = crud.CreateCustomers(uow, cust)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Customer added successfully")

	// customers := crud.GetAllCustomers(uow)
	// for i, c := range customers {
	// 	fmt.Println("Customer", i, ":", c)
	// }

	cust := customer.Customer{}
	cust.ID = 102

	// err = crud.UpdateCustomerName(uow, cust, "Ben")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Name successfully updated")

	// currentCustomer := crud.GetCustomers(uow, cust)
	// fmt.Println("Customer:", currentCustomer)

	err = crud.DeleteCustomer(uow, cust)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Customer Deleted")

}

package main

import (
	"fmt"
	customer "gormTest/customerrepo/model"
	"gormTest/customerrepo/repository"
	"gormTest/customerrepo/unitofwork"

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

	uow := unitofwork.NewUnitOfWork(db, false)

	repo := repository.NewRepository()
	// Adding new customers
	cust := customer.Customer{
		Name:    "Tim",
		Address: "Mumbai",
		Age:     21,
	}
	cust.ID = 101

	uow.Committed = true
	if err := repo.AddCustomer(uow, cust); err != nil {
		uow.Complete()
		return
	}
	uow.Commit()
	fmt.Println("Customer added")

	// Get All Customers from the table
	// cust := &customer.Customer{}
	// if err := repo.Get(uow, cust); err != nil {
	// 	uow.Complete()
	// 	return
	// }
	// uow.Commit()
	// fmt.Println(cust)
	// for i, c := range customers {
	// 	fmt.Println("Customer", i, ":", c)
	// }

	// Updating Customer Name
	// cust := customer.Customer{}
	// cust.ID = 104

	// newCust := customer.Customer{
	// 	Name: "Tom",
	// }

	// if err := repo.UpdateCustomer(uow, cust, newCust); err != nil {
	// 	uow.Complete()
	// 	return
	// }
	// uow.Commit()
	// fmt.Println("Customer successfully updated")

	// Delete customer from table
	// if err := repo.DeleteCustomer(uow, cust); err != nil {
	// 	uow.Complete()
	// 	return
	// }
	// uow.Commit()
	// fmt.Println("Customer Deleted")

}

package main

import (
	"fmt"
	"gormTest/customerrepo/model"
	"gormTest/customerrepo/repository"

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

	db.AutoMigrate(&model.Customer{}, &model.Order{})
	// db.Model(&model.Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "RESTRICT")

	uow := repository.NewUnitOfWork(db, false)

	repo := repository.NewRepository()

	// =======================Adding new customers=========================
	cust := &model.Customer{

		Name: "Tom",
		Orders: []model.Order{
			{
				OrderID:   502,
				OrderName: "Order1",
			},
			{
				OrderID:   504,
				OrderName: "Order4",
			},
		},
	}
	cust.ID = 101

	uow.Committed = true
	if err := repo.Update(uow, cust); err != nil {
		uow.Complete()
		return
	}
	uow.Commit()
	fmt.Println("Customer added")

	// ==================Get All Customers from the table======================
	// cust := &model.Customer{}
	// cust.ID = 101
	// if err := repo.Get(uow, cust); err != nil {
	// 	uow.Complete()
	// 	return
	// }
	// uow.Commit()
	// fmt.Println(*cust)

	// ===========================Updating Customer Name=====================
	// cust := &model.Customer{
	// 	Name: "Ben",
	// }
	// cust.ID = 103

	// if err := repo.Update(uow, cust); err != nil {
	// 	uow.Complete()
	// 	return
	// }
	// uow.Commit()
	// fmt.Println("Customer successfully updated")

	// ===================Delete customer from table============================
	// cust := &model.Customer{}
	// cust.ID = 103
	// if err := repo.Delete(uow, cust); err != nil {
	// 	uow.Complete()
	// 	return
	// }
	// uow.Commit()
	// fmt.Println("Customer Deleted")

	// ===================Delete customer from table (Unscoped)============================
	// cust := &model.Customer{}
	// cust.ID = 103
	// if err := repo.UnscopedDelete(uow, cust); err != nil {
	// 	uow.Complete()
	// 	return
	// }
	// uow.Commit()
	// fmt.Println("Customer Deleted")

}

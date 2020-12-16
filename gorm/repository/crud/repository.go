package crud

import (
	"gormTest/repository/customer"
	"gormTest/repository/unitOfWork"
)

func CreateCustomers(uow *unitOfWork.UnitOfWork, cust customer.Customer) error {

	tx := uow.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Debug().Create(&cust).Error; err != nil {
		tx.Rollback()
		return err
	}
	uow.Committed = true

	return tx.Commit().Error
}

func GetAllCustomers(uow *unitOfWork.UnitOfWork) []customer.Customer {
	uow.Committed = false

	cust := []customer.Customer{}

	uow.DB.Debug().Model(&customer.Customer{}).Find(&cust)
	return cust
}

func GetCustomers(uow *unitOfWork.UnitOfWork, cust customer.Customer) customer.Customer {
	uow.Committed = false

	uow.DB.Debug().Model(&customer.Customer{}).First(&cust)
	return cust
}

func UpdateCustomerName(uow *unitOfWork.UnitOfWork, cust customer.Customer, newName string) error {

	uow.Committed = true

	tx := uow.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Debug().Model(&cust).Update("Name", newName).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func DeleteCustomer(uow *unitOfWork.UnitOfWork, cust customer.Customer) error {

	uow.Committed = true

	tx := uow.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Debug().Where("id=?", cust.ID).Delete(&cust).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

package repository

import (
	customer "gormTest/customerrepo/model"
)

type Repository interface {
	Get(uow *UnitOfWork, out interface{}) error
	AddCustomer(uow *UnitOfWork, cust customer.Customer) error
	UpdateCustomer(uow *UnitOfWork, cust customer.Customer, newCust customer.Customer) error
	DeleteCustomer(uow *UnitOfWork, cust customer.Customer) error
}

type GormRepository struct {
}

func NewRepository() *GormRepository {
	return &GormRepository{}
}

func (g *GormRepository) Get(uow *UnitOfWork, out interface{}) error {
	db := uow.DB

	if err := db.Debug().Find(out).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) AddCustomer(uow *UnitOfWork, cust customer.Customer) error {

	if err := uow.DB.Debug().Create(&cust).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) UpdateCustomer(uow *UnitOfWork, cust customer.Customer, newCust customer.Customer) error {

	if err := uow.DB.Debug().Model(&cust).Update(&newCust).Error; err != nil {
		return err
	}
	return nil

}

func (g *GormRepository) DeleteCustomer(uow *UnitOfWork, cust customer.Customer) error {

	if err := uow.DB.Debug().Where("id=?", cust.ID).Delete(&cust).Error; err != nil {
		return err
	}
	return nil
}

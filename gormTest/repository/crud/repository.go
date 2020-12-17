package crud

import (
	customer "gormTest/repository/model"
	"gormTest/repository/unitofwork"
)

type Repository interface {
	Get(uow *unitofwork.UnitOfWork, out interface{}) error
	AddCustomer(uow *unitofwork.UnitOfWork, cust customer.Customer) error
	UpdateCustomer(uow *unitofwork.UnitOfWork, cust customer.Customer, newCust customer.Customer) error
	DeleteCustomer(uow *unitofwork.UnitOfWork, cust customer.Customer) error
}

type GormRepository struct {
}

func NewRepository() *GormRepository {
	return &GormRepository{}
}

func (g *GormRepository) Get(uow *unitofwork.UnitOfWork, out interface{}) error {
	db := uow.DB

	if err := db.Debug().Find(out).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) AddCustomer(uow *unitofwork.UnitOfWork, cust customer.Customer) error {

	if err := uow.DB.Debug().Create(&cust).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) UpdateCustomer(uow *unitofwork.UnitOfWork, cust customer.Customer, newCust customer.Customer) error {

	if err := uow.DB.Debug().Model(&cust).Update(&newCust).Error; err != nil {
		return err
	}
	return nil

}

func (g *GormRepository) DeleteCustomer(uow *unitofwork.UnitOfWork, cust customer.Customer) error {

	if err := uow.DB.Debug().Where("id=?", cust.ID).Delete(&cust).Error; err != nil {
		return err
	}
	return nil
}

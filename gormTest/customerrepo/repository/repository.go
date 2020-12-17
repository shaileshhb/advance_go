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

	if err := db.Debug().Preload("Orders").First(out).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB

	if err := db.Debug().Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB

	if err := db.Debug().Model(entity).Update(entity).Error; err != nil {
		return err
	}
	return nil

}

func (g *GormRepository) Delete(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB

	if err := db.Debug().Delete(entity).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) UnscopedDelete(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB

	if err := db.Debug().Unscoped().Delete(entity).Error; err != nil {
		return err
	}
	return nil
}

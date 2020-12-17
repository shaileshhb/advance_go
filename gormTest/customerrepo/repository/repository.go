package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/advance_go/gormTest/customerrepo/model"
)

type Repository interface {
	Get(uow *UnitOfWork, out interface{}) error
	AddCustomer(uow *UnitOfWork, cust model.Customer) error
	UpdateCustomer(uow *UnitOfWork, cust model.Customer, newCust model.Customer) error
	DeleteCustomer(uow *UnitOfWork, cust model.Customer) error
}

type GormRepository struct{}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

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

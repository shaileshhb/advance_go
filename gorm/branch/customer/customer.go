package customer

import (
	"gormTest/branch/model"

	uuid "github.com/satori/go.uuid"
)

type BankCustomer struct {
	model.Model
	Name   string
	Amount int
	BankID uuid.UUID `gorm:"type:varchar(36)"`
}

package customer

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shaileshhb/advance_go/gormTest/branch/model"
)

type BankCustomer struct {
	model.Model
	Name   string
	Amount int
	BankID uuid.UUID `gorm:"type:varchar(36)"`
}

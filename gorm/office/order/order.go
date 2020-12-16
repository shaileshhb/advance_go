package order

import uuid "github.com/satori/go.uuid"

type Order struct {
	OrderID    int `gorm:"primary_key"`
	OrderName  *string
	CustomerID uuid.UUID
}

package model

type Order struct {
	OrderID    int `gorm:"primary_key"`
	OrderName  string
	CustomerID uint
}

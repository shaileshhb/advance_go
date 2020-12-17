package model

import (
	uuid "github.com/satori/go.uuid"
)

type Model struct {
	ID uuid.UUID `gorm:"primary_key"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt *time.Time `sql:"index"`
}

// // BeforeCreate will set a UUID rather than numeric ID.
// func (base *Model) BeforeCreate(scope *gorm.Scope) error {
// 	uuid := uuid.NewV4()

// 	return scope.SetColumn("ID", uuid)
// }

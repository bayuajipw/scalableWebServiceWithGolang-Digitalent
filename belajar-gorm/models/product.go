package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(100)"`
	Brand     string `gorm:"not null;type:varchar(100)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("PRODUCT before create()")

	if len(p.Name) < 4 {
		err = errors.New("Product name should be more than 4")
	}

	return
}

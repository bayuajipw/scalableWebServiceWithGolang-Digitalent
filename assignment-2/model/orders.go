package model

import "time"

type Orders struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;type:varchar(50)"`
	OrderAt      time.Time
}

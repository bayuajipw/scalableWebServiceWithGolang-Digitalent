package model

type Items struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"not null;type:varchar(10)"`
	Description string `gorm:"not null;type:varchar(50)"`
	Quantity    int    `gorm:"not null;type:bigint"`
	OrderID     int    `gorm:"not null;type:bigint"`
}

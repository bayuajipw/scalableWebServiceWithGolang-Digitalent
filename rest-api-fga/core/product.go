package core

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Quantity    int
}

type ProductImpl interface {
	GetAllProducts() ([]Product, error)
}

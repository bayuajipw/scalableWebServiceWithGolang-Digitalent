package product

import (
	"rest-api-fga/core"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllProducts() ([]core.Product, error) {
	products := []core.Product{}

	err := r.db.Find(&products).Error

	return products, err
}

package repository

import (
	"final_project/model"
)

func (db *Database) GetProducts(pagination *model.Pagination) ([]*model.Product, error) {
	var products []*model.Product
	err := db.DB.Preload("Variants").Find(&products).Error
	return products, err
}

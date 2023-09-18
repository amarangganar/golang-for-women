package repository

import (
	"final_project/model"
)

type ProductList struct {
	Data       []*model.Product
	Pagination *model.Paginator
}

func (db *Database) GetProducts(param *model.ListQueryParam) (*ProductList, error) {
	var products []*model.Product

	paginator := Paginator(db, param, products, "name LIKE ?")

	err := db.DB.Preload("Admin").Preload("Variants").Scopes(Paginate(paginator.Limit, paginator.Offset)).Where("name LIKE ?", "%"+param.Search+"%").Find(&products).Error

	result := &ProductList{
		Data:       products,
		Pagination: paginator,
	}

	return result, err
}

func (db *Database) CreateProduct(body *model.NewProduct) (*model.Product, error) {
	product := model.Product{
		Name:     body.Name,
		ImageURL: body.File,
		AdminID:  *body.AdminID,
	}

	if err := db.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	err := db.DB.Preload("Admin").First(&product).Error

	return &product, err
}

func (db *Database) GetProduct(uuid string) (*model.Product, error) {
	var product *model.Product

	err := db.DB.Preload("Admin").Preload("Variants").Where("uuid = ?", uuid).First(&product).Error

	return product, err
}

func (db *Database) UpdateProduct(body *model.ExistingProduct) (*model.Product, error) {
	product, err := db.GetProduct(body.UUID)
	if err != nil {
		return nil, err
	}

	product.Name = body.Name
	if body.File != "" {
		product.ImageURL = body.File
	}

	err = db.DB.Save(&product).Error

	return product, err
}

func (db *Database) DeleteProduct(uuid string) error {
	var product *model.Product

	err := db.DB.Where("uuid = ?", uuid).First(&product).Error
	if err != nil {
		return err
	}

	err = db.DB.Delete(&product).Error

	return err
}

package repository

import (
	"final_project/model"

	"github.com/google/uuid"
)

type VariantList struct {
	Data       []*model.Variant
	Pagination *model.Paginator
}

func (db *Database) GetVariants(param *model.ListQueryParam) (*VariantList, error) {
	var variants []*model.Variant

	paginator := Paginator(db, param, variants, "variant_name LIKE ?")

	err := db.DB.Preload("Product.Admin").Scopes(Paginate(paginator.Limit, paginator.Offset)).Where("variant_name LIKE ?", "%"+param.Search+"%").Find(&variants).Error

	result := &VariantList{
		Data:       variants,
		Pagination: paginator,
	}

	return result, err
}

func (db *Database) CreateVariant(body *model.NewVariant) (*model.Variant, error) {
	product, err := db.GetProduct(body.ProductID)

	if err != nil {
		return nil, err
	}

	variant := model.Variant{
		UUID:        uuid.NewString(),
		VariantName: body.VariantName,
		Quantity:    *body.Quantity,
		ProductID:   product.ID,
	}

	err = db.DB.Create(&variant).Error

	if err != nil {
		return nil, err
	}

	err = db.DB.Preload("Product.Admin").First(&variant).Error

	return &variant, err
}

func (db *Database) GetVariant(uuid string) (*model.Variant, error) {
	var variant *model.Variant

	err := db.DB.Preload("Product.Admin").Where("uuid = ?", uuid).First(&variant).Error

	return variant, err
}

func (db *Database) UpdateVariant(uuid string, body *model.ExistingVariant) (*model.Variant, error) {
	variant, err := db.GetVariant(uuid)
	if err != nil {
		return nil, err
	}

	variant.VariantName = body.VariantName
	variant.Quantity = *body.Quantity

	err = db.DB.Save(&variant).Error

	if body.ProductID != "" {
		product, err := db.GetProduct(body.ProductID)
		if err != nil {
			return nil, err
		}

		db.DB.Model(&variant).Association("Product").Replace(product)
	}

	return variant, err
}

func (db *Database) DeleteVariant(uuid string) error {
	var variant *model.Variant

	err := db.DB.Where("uuid = ?", uuid).First(&variant).Error
	if err != nil {
		return err
	}

	err = db.DB.Delete(&variant).Error

	return err
}

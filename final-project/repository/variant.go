package repository

import (
	"final_project/model"
)

type VariantList struct {
	Data       []*model.Variant
	Pagination *model.Paginator
}

func (db *Database) GetVariants(param *model.ListQueryParam) (*VariantList, error) {
	var variants []*model.Variant

	paginator := Paginator(db, param, variants, "variant_name = ?")

	err := db.DB.Preload("Product").Scopes(Paginate(paginator.Limit, paginator.Offset)).Find(&variants).Where("variant_name = ?", param.Search).Error

	result := &VariantList{
		Data:       variants,
		Pagination: paginator,
	}

	return result, err
}

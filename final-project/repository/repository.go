package repository

import (
	"final_project/model"
	"math"

	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func Paginator(db *Database, param *model.ListQueryParam, m interface{}, search_by string) *model.Paginator {
	limit := 10
	if param.Limit != nil {
		limit = *param.Limit
	}

	offset := param.Offset

	var count int64
	db.DB.Model(&m).Where(search_by, "%"+param.Search+"%").Count(&count)

	var last_page int
	if count > 0 {
		last_page = int(math.Ceil(float64(count) / float64(limit)))
	} else {
		last_page = 1
	}

	paginator := &model.Paginator{
		Limit:    limit,
		Offset:   offset,
		Page:     offset + 1,
		LastPage: last_page,
		Total:    count,
	}

	return paginator
}

func Paginate(limit int, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

package api

import (
	"final_project/model"
	"final_project/repository"
	v "final_project/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VariantDatabase interface {
	GetVariants(param *model.ListQueryParam) (*repository.VariantList, error)
	CreateVariant(body *model.NewVariant) (*model.Variant, error)
	GetVariant(uuid string) (*model.Variant, error)
	UpdateVariant(uuid string, body *model.ExistingVariant) (*model.Variant, error)
	DeleteVariant(uuid string) error
}

type VariantAPI struct {
	DB VariantDatabase
}

func (db *VariantAPI) GetVariants(ctx *gin.Context) {
	param := model.ListQueryParam{}

	if err := ctx.ShouldBind(&param); err != nil {
		v.Validate(ctx, &param, err, "Fetch variants failed.", v.Paginate)
		return
	}

	result, err := db.DB.GetVariants(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Fetch products failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Variants fetched.",
		"data":       result.Data,
		"pagination": result.Pagination,
	})
}

func (db *VariantAPI) CreateVariant(ctx *gin.Context) {
	body := model.NewVariant{}

	if err := ctx.ShouldBind(&body); err != nil {
		v.Validate(ctx, &body, err, "create variant failed", v.ValidateField)
		return
	}

	res, err := db.DB.CreateVariant(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "create variant failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "new variant created.",
		"data":    res,
	})
}

func (api *VariantAPI) GetVariant(ctx *gin.Context) {
	body := model.UUIDVariant{}

	if err := ctx.ShouldBindUri(&body); err != nil {
		v.Validate(ctx, &body, err, "fetch variant failed.", v.ValidateField)
		return
	}

	res, err := api.DB.GetVariant(body.UUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fetch variant failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Variant fetched.",
		"data":    res,
	})
}

func (api *VariantAPI) UpdateVariant(ctx *gin.Context) {
	uri := model.UUIDVariant{}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		v.Validate(ctx, &uri, err, "update variant failed", v.ValidateField)
		return
	}

	body := model.ExistingVariant{}
	if err := ctx.ShouldBind(&body); err != nil {
		v.Validate(ctx, &body, err, "update variant failed", v.ValidateField)
		return
	}

	res, err := api.DB.UpdateVariant(uri.UUID, &body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "update variant failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "variant updated",
		"data":    res,
	})
}

func (api *VariantAPI) DeleteVariant(ctx *gin.Context) {
	uri := model.UUIDVariant{}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		v.Validate(ctx, &uri, err, "delete variant failed", v.ValidateField)
		return
	}

	err := api.DB.DeleteVariant(uri.UUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "delete variant failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "variant deleted",
		"data":    nil,
	})
}

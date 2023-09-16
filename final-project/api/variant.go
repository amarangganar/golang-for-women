package api

import (
	"final_project/model"
	v "final_project/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVariants(ctx *gin.Context) {
	param := model.Pagination{}

	if err := ctx.ShouldBind(&param); err != nil {
		v.Validate(ctx, &param, err, "Fetch variants failed.", v.Paginate)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Variants fetched.",
		"data":    param,
	})
}

func CreateVariant(ctx *gin.Context) {
	body := model.NewVariant{}

	if err := ctx.ShouldBind(&body); err != nil {
		v.Validate(ctx, &body, err, "Create variant failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "New variant created.",
		"data":    body,
	})
}

func GetVariant(ctx *gin.Context) {
	body := model.UUIDVariant{}

	if err := ctx.ShouldBindUri(&body); err != nil {
		v.Validate(ctx, &body, err, "Fetch variant failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Variant fetched.",
		"data":    body,
	})
}

func UpdateVariant(ctx *gin.Context) {
	uri := model.UUIDVariant{}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		v.Validate(ctx, &uri, err, "Update variant failed.", v.ValidateField)
		return
	}

	body := model.ExistingVariant{}
	if err := ctx.ShouldBind(&body); err != nil {
		v.Validate(ctx, &body, err, "Update variant failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Variant updated.",
		"data":    body,
	})
}

func DeleteVariant(ctx *gin.Context) {
	body := model.UUIDVariant{}

	if err := ctx.ShouldBindUri(&body); err != nil {
		v.Validate(ctx, &body, err, "Delete variant failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Variant deleted.",
		"data":    body,
	})
}

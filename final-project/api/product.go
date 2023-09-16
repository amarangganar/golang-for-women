package api

import (
	"final_project/model"
	v "final_project/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetProducts(ctx *gin.Context) {
	param := model.Pagination{}

	if err := ctx.ShouldBind(&param); err != nil {
		v.Validate(ctx, &param, err, "Fetch products failed.", v.Paginate)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Products fetched.",
		"data":    param,
	})
}

func CreateProduct(ctx *gin.Context) {
	body := model.NewProduct{}
	body.Name = ctx.Request.FormValue("name")

	// check file
	file, errFile := ctx.FormFile("file")
	if errFile == nil {
		body.File = file.Filename
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		v.Validate(ctx, &body, err, "Create product failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "New product created.",
		"data":    body,
	})
}

func GetProduct(ctx *gin.Context) {
	body := model.UUIDProduct{}

	if err := ctx.ShouldBindUri(&body); err != nil {
		v.Validate(ctx, &body, err, "Fetch product failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product fetched.",
		"data":    body,
	})
}

func UpdateProduct(ctx *gin.Context) {
	body := model.ExistingProduct{}

	body.UUID = ctx.Param("uuid")
	body.Name = ctx.Request.FormValue("name")

	// check file
	file, errFile := ctx.FormFile("file")
	if errFile == nil {
		body.File = file.Filename
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		v.Validate(ctx, &body, err, "Update product failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated.",
		"data":    body,
	})
}

func DeleteProduct(ctx *gin.Context) {
	body := model.UUIDProduct{}

	if err := ctx.ShouldBindUri(&body); err != nil {
		v.Validate(ctx, &body, err, "Delete product failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted.",
		"data":    body,
	})
}

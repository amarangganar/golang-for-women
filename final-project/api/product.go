package api

import (
	"final_project/model"
	"final_project/repository"
	v "final_project/validator"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type ProductDatabase interface {
	GetProducts(param *model.ListQueryParam) (*repository.ProductList, error)
	CreateProduct(body *model.NewProduct) (*model.Product, error)
	GetProduct(uuid string) (*model.Product, error)
	UpdateProduct(uuid string, body *model.ExistingProduct) (*model.Product, error)
	DeleteProduct(uuid string) error
}

type ProductAPI struct {
	DB         ProductDatabase
	Cloudinary *cloudinary.Cloudinary
}

func (api *ProductAPI) GetProducts(ctx *gin.Context) {
	param := model.ListQueryParam{}

	if err := ctx.ShouldBind(&param); err != nil {
		v.Validate(ctx, &param, err, "fetch products failed", v.Paginate)
		return
	}

	result, err := api.DB.GetProducts(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fetch products failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "products fetched",
		"data":       result.Data,
		"pagination": result.Pagination,
	})
}

func (api *ProductAPI) CreateProduct(ctx *gin.Context) {
	body := model.NewProduct{}
	body.Name = ctx.Request.FormValue("name")

	// check file
	file, file_meta, errFile := ctx.Request.FormFile("file")
	if errFile == nil {
		body.File = file_meta.Filename
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		v.Validate(ctx, &body, err, "create product failed", v.ValidateField)
		return
	}

	// upload file to cloudinary
	upload, err := api.Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "products",
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "create product failed",
			"error":   err.Error(),
		})
	}
	// set image url from cloudinary
	body.File = upload.URL

	// assign admin who created the product
	admin := ctx.MustGet("admin").(jwt.MapClaims)
	admin_id := uint(admin["id"].(float64))
	body.AdminID = &admin_id

	result, err := api.DB.CreateProduct(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "create product failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "New product created.",
		"data":    result,
	})
}

func (api *ProductAPI) GetProduct(ctx *gin.Context) {
	body := model.UUIDProduct{}

	if err := ctx.ShouldBindUri(&body); err != nil {
		v.Validate(ctx, &body, err, "fetch product failed", v.ValidateField)
		return
	}

	result, err := api.DB.GetProduct(body.UUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fetch product failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product fetched",
		"data":    result,
	})
}

func (api *ProductAPI) UpdateProduct(ctx *gin.Context) {
	uri := model.UUIDProduct{}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		v.Validate(ctx, &uri, err, "update variant failed", v.ValidateField)
		return
	}

	body := model.ExistingProduct{}
	body.Name = ctx.Request.FormValue("name")

	// check file
	file, file_meta, errFile := ctx.Request.FormFile("file")
	if errFile == nil {
		body.File = file_meta.Filename
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		v.Validate(ctx, &body, err, "update product failed", v.ValidateField)
		return
	}

	if body.File != "" {
		// upload file to cloudinary
		upload, err := api.Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
			Folder: "products",
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "update product failed",
				"error":   err.Error(),
			})
		}
		// set image url from cloudinary
		body.File = upload.URL
	}

	result, err := api.DB.UpdateProduct(uri.UUID, &body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "update product failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product updated",
		"data":    result,
	})
}

func (api *ProductAPI) DeleteProduct(ctx *gin.Context) {
	uri := model.UUIDProduct{}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		v.Validate(ctx, &uri, err, "delete product failed", v.ValidateField)
		return
	}

	err := api.DB.DeleteProduct(uri.UUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "delete product failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted.",
		"data":    nil,
	})
}

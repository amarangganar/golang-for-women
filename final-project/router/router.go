package router

import (
	"final_project/api"
	authentication "final_project/auth"
	"final_project/repository"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
)

func Execute(db *repository.Database, cld *cloudinary.Cloudinary) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authHandler := authentication.Auth{DB: db}
	productHandler := api.ProductAPI{DB: db, Cloudinary: cld}
	variantHandler := api.VariantAPI{DB: db}

	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.SignIn)
		auth.POST("/register", authHandler.Register)
	}

	// Product
	product := r.Group("/products")
	{
		product.GET("/", productHandler.GetProducts)
		product.GET("/:uuid", productHandler.GetProduct)

		product.Use(Authenticate())
		product.POST("/", productHandler.CreateProduct)

		product.Use(ProductAuthorization(db))
		product.PUT("/:uuid", productHandler.UpdateProduct)
		product.DELETE("/:uuid", productHandler.DeleteProduct)
	}

	// Product
	variant := r.Group("/variants")
	{

		variant.GET("/", variantHandler.GetVariants)
		variant.POST("/", api.CreateVariant)
		variant.GET("/:uuid", api.GetVariant)
		variant.PUT("/:uuid", api.UpdateVariant)
		variant.DELETE("/:uuid", api.DeleteVariant)
	}

	return r
}

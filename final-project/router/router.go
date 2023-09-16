package router

import (
	"final_project/api"
	authentication "final_project/auth"
	"os"

	"github.com/gin-gonic/gin"
)

func Execute() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/login", authentication.SignIn)
		auth.POST("/register", authentication.Register)
	}

	// Product
	product := r.Group("/products")
	{
		product.GET("/", api.GetProducts)
		product.POST("/", api.CreateProduct)
		product.GET("/:uuid", api.GetProduct)
		product.PUT("/:uuid", api.UpdateProduct)
		product.DELETE("/:uuid", api.DeleteProduct)
	}

	// Product
	variant := r.Group("/variants")
	{

		variant.GET("/", api.GetVariants)
		variant.POST("/", api.CreateVariant)
		variant.GET("/:uuid", api.GetVariant)
		variant.PUT("/:uuid", api.UpdateVariant)
		variant.DELETE("/:uuid", api.DeleteVariant)
	}

	r.Run(os.Getenv("SERVER_PORT"))
}

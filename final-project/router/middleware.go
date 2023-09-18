package router

import (
	"final_project/auth"
	"final_project/model"
	"final_project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := auth.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
				"error":   err.Error(),
			})
			return
		}

		ctx.Set("admin", verifyToken)
		ctx.Next()
	}
}

func ProductAuthorization(db *repository.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")

		admin := ctx.MustGet("admin").(jwt.MapClaims)
		admin_id := uint(admin["id"].(float64))

		var product *model.Product
		err := db.DB.Select("admin_id").Where("uuid = ?", uuid).First(&product).Error
		if err != nil || product.AdminID != admin_id {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "unauthorized",
				"message": "you are not allowed to access this api",
			})

			return
		}

		ctx.Next()
	}
}

func VariantAuthorization(db *repository.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		admin := ctx.MustGet("admin").(jwt.MapClaims)
		admin_id := uint(admin["id"].(float64))

		// authorize create api
		if ctx.Request.Method == "POST" {
			product_id := ctx.Request.FormValue("product_id")

			var product *model.Product
			err := db.DB.Select("admin_id").Where("uuid = ?", product_id).First(&product).Error

			if err != nil || product.AdminID != admin_id {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error":   "unauthorized",
					"message": "you are not the owner of the product",
				})

				return
			}
		}

		// authorize updte and delete
		uuid := ctx.Param("uuid")
		if uuid != "" {
			var variant *model.Variant
			err := db.DB.Where("uuid = ?", uuid).Preload("Product").First(&variant).Error
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error":   "unauthorized",
					"message": "you are not allowed to access this api",
				})
				return
			}

			if ctx.Request.Method == "PUT" {
				product_id := ctx.Request.FormValue("product_id")
				if product_id != "" {
					var product *model.Product
					err := db.DB.Select("admin_id").Where("uuid = ?", product_id).First(&product).Error
					if err != nil || product.AdminID != admin_id {
						ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
							"error":   "unauthorized",
							"message": "you are not the owner of the product",
						})

						return
					}
				}
			}

			if ctx.Request.Method == "DELETE" {
				if variant.Product.ID != admin_id {
					ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
						"error":   "unauthorized",
						"message": "you are not allowed to access this api",
					})
					return
				}
			}
		}

		ctx.Next()
	}
}

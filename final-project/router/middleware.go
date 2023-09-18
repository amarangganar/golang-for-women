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
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "fetch product failed",
				"error":   err.Error(),
			})

			return
		}

		if product.AdminID != admin_id {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "you are not allowed to access this product",
			})
			return
		}

		ctx.Next()
	}
}

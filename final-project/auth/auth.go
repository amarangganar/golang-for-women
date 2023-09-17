package auth

import (
	"final_project/model"
	v "final_project/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	admin := model.AdminSignIn{}

	if err := ctx.ShouldBind(&admin); err != nil {
		v.Validate(ctx, &admin, err, "Sign in failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sign in successfull.",
		"data":    admin,
	})
}

func Register(ctx *gin.Context) {
	admin := model.AdminRegister{}

	if err := ctx.ShouldBind(&admin); err != nil {
		v.Validate(ctx, &admin, err, "Register failed.", v.ValidateField)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register successfull.",
		"data":    admin,
	})
}

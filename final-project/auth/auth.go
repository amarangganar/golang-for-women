package auth

import (
	"errors"
	"final_project/helpers"
	"final_project/model"
	v "final_project/validator"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthDatabase interface {
	SignIn(admin *model.AdminSignIn) (*model.Admin, error)
	Register(new_admin *model.AdminRegister) (*model.Admin, error)
}

type Auth struct {
	DB AuthDatabase
}

func (db *Auth) SignIn(ctx *gin.Context) {
	admin := model.AdminSignIn{}

	if err := ctx.ShouldBind(&admin); err != nil {
		v.Validate(ctx, &admin, err, "sign in failed.", v.ValidateField)
		return
	}

	res, err := db.DB.SignIn(&admin)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Println(helpers.ComparePassword([]byte(res.Password), []byte(admin.Password)))

	if !helpers.ComparePassword([]byte(res.Password), []byte(admin.Password)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid password",
			"error":   nil,
		})
		return
	}

	token := generateToken(res.ID, res.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "sign in successfull",
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

func (db *Auth) Register(ctx *gin.Context) {
	admin := model.AdminRegister{}

	if err := ctx.ShouldBind(&admin); err != nil {
		v.Validate(ctx, &admin, err, "register failed", v.ValidateField)
		return
	}

	result, err := db.DB.Register(&admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "register failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "register successfull",
		"data":    result,
	})
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errors.New("you need to sign in")
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("you need to sign in")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("you need to sign in")
	}

	expClaim, exists := claims["expired_at"]
	if !exists {
		return nil, errors.New("expire claim is missing")
	}

	expStr, ok := expClaim.(string)
	if !ok {
		return nil, errors.New("expire claim is not a valid type")
	}

	expTime, err := time.Parse(time.RFC3339, expStr)
	if err != nil {
		return nil, errors.New("error parsing expiration time")
	}

	if time.Now().After(expTime) {
		return nil, errors.New("token is expired")
	}

	return token.Claims.(jwt.MapClaims), nil
}

func generateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":         id,
		"email":      email,
		"expired_at": time.Now().Add(time.Minute * 120),
	}

	parse_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed_token, err := parse_token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return err.Error()
	}

	return signed_token
}

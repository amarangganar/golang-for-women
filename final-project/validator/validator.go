package validator

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type validate func(obj interface{}, err error) map[string]interface{}

func Validate(ctx *gin.Context, param interface{}, err error, msg string, v validate) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		error_field := v(param, err)

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": msg,
			"error":   error_field,
		})

		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal server error.",
		"error":   err.Error(),
	})
}

func Paginate(obj interface{}, err error) map[string]interface{} {
	error_field := make(map[string]interface{})

	for _, err := range err.(validator.ValidationErrors) {
		var message string

		if err.Tag() == "number" {
			message = fmt.Sprintf("%s should be a %s.", err.Field(), err.Tag())
		} else if err.Tag() == "min" {
			message = fmt.Sprintf("%s should be at least %s.", err.Field(), err.Param())
		}

		field, _ := reflect.ValueOf(obj).Elem().Type().FieldByName(err.Field())
		fieldName, _ := field.Tag.Lookup("form")

		// Get the field of the slice element that we want to set.
		error_field[fieldName] = message
	}

	return error_field
}

func ValidateField(obj interface{}, err error) map[string]interface{} {
	error_field := make(map[string]interface{})

	for _, err := range err.(validator.ValidationErrors) {
		var message string

		if err.Tag() == "required" {
			message = fmt.Sprintf("%s is %s.", err.Field(), err.Tag())
		} else if err.Tag() == "uuid" {
			message = fmt.Sprintf("%s should be a valid %s.", err.Field(), err.Tag())
		} else if err.Tag() == "email" {
			message = fmt.Sprintf("%s should be a valid %s.", err.Field(), err.Tag())
		} else if err.Tag() == "min" && err.Kind().String() == "string" {
			message = fmt.Sprintf("%s should be at least %s characters.", err.Field(), err.Param())
		}

		field, _ := reflect.ValueOf(obj).Elem().Type().FieldByName(err.Field())
		fieldName, is_form := field.Tag.Lookup("form")
		if !is_form {
			fieldName, _ := field.Tag.Lookup("uri")
			error_field[fieldName] = message
		} else {
			error_field[fieldName] = message
		}

	}

	return error_field
}

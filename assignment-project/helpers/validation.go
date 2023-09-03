package helpers

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

func Validate(field interface{}) map[string]interface{} {
	var error map[string]interface{}

	_, validateErr := govalidator.ValidateStruct(field)
	if validateErr != nil {
		error = getErrorField(validateErr)
	} else {
		error = nil
	}

	return error
}

// this only works for nested struct two levels
// also, doesn't work to get error field of each struct on an array of struct
func getErrorField(err error) map[string]interface{} {
	errorField := make(map[string]interface{})
	errs := err.(govalidator.Errors).Errors()

	for _, e := range errs {
		casted, ok := e.(govalidator.Error)
		if ok {
			errorField[casted.Name] = casted.Error()
		} else {
			nested := e.(govalidator.Errors).Errors()

			// this is hacky as hell
			// since the key is the key of the parent object
			// we assign the field name using the key object on the index 0
			// and we can't know on which index the error is happening
			// because the error only tells us which attribute has an error, not the whole field.
			nestedCasted := nested[0].(govalidator.Error)
			errorField[nestedCasted.Name] = strings.Split(e.Error(), ";")
		}
	}

	return errorField
}

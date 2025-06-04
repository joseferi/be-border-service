package validators

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error, input any) []map[string]string {
	var errors []map[string]string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		val := reflect.TypeOf(input)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		for _, e := range validationErrors {
			// Ambil nama field dari struct
			if field, found := val.FieldByName(e.StructField()); found {
				// Ambil nama json tag
				jsonTag := field.Tag.Get("json")
				if jsonTag == "" || jsonTag == "-" {
					jsonTag = e.Field()
				}
				// Tambahkan ke daftar error
				errors = append(errors, map[string]string{
					jsonTag: jsonTag + " " + e.Tag(),
				})
			}
		}
	}
	return errors
}

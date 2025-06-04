package databasex

import (
	"fmt"
	"reflect"
	"strings"
)

func BuildWhereClause(filter interface{}) (string, []interface{}) {
	v := reflect.ValueOf(filter)
	t := reflect.TypeOf(filter)

	var conditions []string
	var args []interface{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		dbTag := field.Tag.Get("db")
		if dbTag == "" {
			continue
		}

		// Cek nilai non-zero
		if !value.IsZero() {
			conditions = append(conditions, fmt.Sprintf("%s = ?", dbTag))
			args = append(args, value.Interface())
		}
	}

	if len(conditions) == 0 {
		return "", nil
	}

	return " WHERE " + strings.Join(conditions, " AND "), args
}

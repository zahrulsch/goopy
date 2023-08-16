package ts

import (
	"reflect"

	"golang.org/x/exp/slices"
)

func FilterPrimitive(target interface{}) (res string) {
	types := reflect.TypeOf(target)
	typesKind := types.Kind()
	res = "unknown"

	if slices.Contains(numbersKind, typesKind) {
		res = "number"
		return
	}

	if typesKind == reflect.String {
		res = "string"
		return
	}

	if typesKind == reflect.Bool {
		res = "boolean"
		return
	}

	return
}

package ts

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/slices"
)

func FilterMap(target interface{}) (res string) {
	res = "unknown"
	keyAs := "string"
	valueAs := "unknown"

	types := reflect.TypeOf(target)
	key := types.Key()
	value := types.Elem()

	if slices.Contains(numbersKind, key.Kind()) {
		keyAs = "number"
	}

	if value.Kind() == reflect.Struct {
		if str := FilterStruct(reflect.Indirect(reflect.New(value)).Interface()); str != res {
			res = fmt.Sprintf("{ [key: %v]: %v }", keyAs, str)
			return
		}
	}

	if prim := FilterPrimitive(reflect.Indirect(reflect.New(value)).Interface()); prim != res {
		res = fmt.Sprintf("{ [key: %v]: %v }", keyAs, prim)
		return
	}

	if ptr := FilterPointer(reflect.Indirect(reflect.New(value)).Interface()); ptr != res {
		res = fmt.Sprintf("{ [key: %v]: %v }", keyAs, ptr)
		return
	}

	res = fmt.Sprintf("{ [key: %v]: %v }", keyAs, valueAs)

	return
}

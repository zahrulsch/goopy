package ts

import (
	"fmt"
	"reflect"
)

func FilterSlice(target interface{}) (res string) {
	defaultValueType := "unknown"

	elem := reflect.Indirect(reflect.New(reflect.TypeOf(target).Elem()))

	if elem.Kind() == reflect.Struct {
		if str := FilterStruct(elem.Interface()); str != defaultValueType {
			res = fmt.Sprintf("Array<%v>", str)
			return
		}
	}

	if prim := FilterPrimitive(elem.Interface()); prim != defaultValueType {
		res = fmt.Sprintf("Array<%v>", prim)
		return
	}

	if ptr := FilterPointer(elem.Interface()); ptr != defaultValueType {
		res = fmt.Sprintf("Array<%v>", ptr)
		return
	}

	res = fmt.Sprintf("Array<%v>", defaultValueType)
	return
}

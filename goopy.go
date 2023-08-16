package goopy

import (
	"reflect"

	"github.com/zahrulsch/goopy/ts"
)

func GenTS(target interface{}) (res string) {
	types := reflect.TypeOf(target)
	typesKind := types.Kind()
	res = "unknown"

	if prim := ts.FilterPrimitive(target); prim != "unknown" {
		res = prim
		return
	}

	if typesKind == reflect.Ptr {
		res = ts.FilterPointer(target)
		return
	}

	if typesKind == reflect.Struct {
		res = ts.FilterStruct(target)
		return
	}

	if typesKind == reflect.Slice {
		res = ts.FilterSlice(target)
		return
	}

	return
}

func GenJS(target interface{}) (res string) {
	return
}

// hilangkan
var numbersKind = []reflect.Kind{
	reflect.Int,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
}

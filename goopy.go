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

	if typesKind == reflect.Map {
		res = ts.FilterMap(target)
		return
	}

	return
}

func GenJS(target interface{}) (res string) {
	return
}

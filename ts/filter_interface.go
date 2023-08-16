package ts

import (
	"fmt"
	"reflect"
)

func FilterInterface(target reflect.Value) (res string) {
	res = "null"
	if !target.IsNil() {
		kind := target.Elem().Kind()

		if kind == reflect.Struct {
			if str := FilterStruct(target.Elem().Interface()); str != "unknown" {
				res = str
				return
			}
		}

		if kind == reflect.Ptr {
			if ptr := FilterPointer(target.Elem().Interface()); ptr != "unknown" {
				res = ptr
				return
			}
		}

		if kind == reflect.Map {
			if m := FilterMap(target.Elem().Interface()); m != "unknown" {
				res = m
				return
			}
		}

		if kind == reflect.Slice {
			if slice := FilterSlice(target.Elem().Interface()); slice != "unknown" {
				res = slice
				return
			}
		}

		if prim := FilterPrimitive(target.Elem().Interface()); prim != "unknown" {
			res = prim
			return
		}

		fmt.Println(kind)

	}

	return
}

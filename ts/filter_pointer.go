package ts

import (
	"reflect"
)

func FilterPointer(target interface{}) (res string) {
	res = "unknown"
	types := reflect.TypeOf(target)
	ntypes := reflect.Indirect(reflect.New(types.Elem()))

	if prim := FilterPrimitive(ntypes.Interface()); prim != res {
		res = prim + " | undefined"
		return
	}

	if ntypes.Kind() == reflect.Slice {
		if slice := FilterSlice(ntypes.Interface()); slice != res {
			res = slice + " | undefined"
			return
		}
	}

	if ntypes.Kind() == reflect.Map {
		if m := FilterMap(ntypes.Interface()); m != res {
			res = m + " | undefined"
			return
		}
	}

	if ntypes.Kind() == reflect.Ptr {
		if ptr := FilterPointer(ntypes.Interface()); ptr != res {
			res = ptr + " | undefined"
			return
		}
	}

	if str := FilterStruct(ntypes.Interface()); str != res {
		res = str + " | undefined"
		return
	}

	return
}

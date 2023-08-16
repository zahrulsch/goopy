package ts

import (
	"fmt"
	"reflect"
	"strings"
)

func FilterStruct(target interface{}) (res string) {
	res = "unknown"
	types := reflect.TypeOf(target)
	values := reflect.ValueOf(target)

	lines := []string{}

	for i := 0; i < values.NumField(); i++ {
		key := types.Field(i).Name
		value := "unknown"

		if tag := types.Field(i).Tag.Get("json"); tag != "" {
			key = tag
		}

		if values.Field(i).Kind() == reflect.Ptr {
			if ptr := FilterPointer(values.Field(i).Interface()); ptr != value {
				line := fmt.Sprintf("%v: %v", key, ptr)
				lines = append(lines, line)
				continue
			}
		}

		if prim := FilterPrimitive(values.Field(i).Interface()); prim != value {
			line := fmt.Sprintf("%v: %v", key, prim)
			lines = append(lines, line)
			continue
		}

		if str := FilterStruct(values.Field(i).Interface()); str != value {
			line := fmt.Sprintf("%v: %v", key, str)
			lines = append(lines, line)

			continue
		}

		fmt.Println(key, value)
	}

	res = fmt.Sprintf("{ %v }", strings.Join(lines, "; "))

	return
}

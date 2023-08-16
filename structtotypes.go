package goopy

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/exp/slices"
)

func primitives(ty reflect.Type, kind reflect.Kind, withUndefined bool) string {
	if kind == reflect.Invalid {
		return "any"
	}

	if kind == reflect.String {
		return "string"
	}

	if kind == reflect.Bool {
		return "boolean"
	}

	if kind == reflect.Ptr {
		elKind := ty.Elem().Kind()

		if converted := primitives(ty, elKind, withUndefined); converted != "" {
			undefined := " | undefined"
			if !withUndefined {
				undefined = ""
			}

			return fmt.Sprintf("%v%v", converted, undefined)
		}
	}

	if slices.Contains(numbersKind, kind) {
		return "number"
	}

	return ""
}

func StructToTYPES(target interface{}) (additionalTypes []string, res string) {
	val := reflect.ValueOf(target)
	kind := val.Kind()

	if prim := primitives(reflect.TypeOf(target), kind, true); prim != "" {
		return nil, prim
	}

	if kind == reflect.Map {
		fmt.Println(kind)
	}

	if kind == reflect.Slice {
		ty := reflect.TypeOf(target).Elem()
		elemKind := ty.Kind()

		if prim := primitives(ty, elemKind, true); prim != "" {
			return nil, fmt.Sprintf("Array<%v>", prim)
		}
	}

	if kind == reflect.Ptr {
		add, res := StructToTYPES(reflect.Indirect(val).Interface())
		additionalTypes = append(additionalTypes, add...)
		return additionalTypes, fmt.Sprintf("%v | undefined", res)
	}

	if kind == reflect.Struct {
		ty := reflect.TypeOf(target)
		lines := []string{}

		for i := 0; i < val.NumField(); i++ {
			key := ty.Field(i).Name
			val := val.Field(i)
			valKind := val.Kind()
			if tag := ty.Field(i).Tag.Get("json"); tag != "" {
				key = tag
			}

			if alterTag := ty.Field(i).Tag.Get("alter_type"); alterTag != "" {
				line := fmt.Sprintf("%v: %v", key, alterTag)
				lines = append(lines, line)
				continue
			}

			if prim := primitives(val.Type(), valKind, true); prim != "" {
				line := fmt.Sprintf("%v: %v", key, prim)
				lines = append(lines, line)
				continue
			}

			if valKind == reflect.Struct {
				add, res := StructToTYPES(val.Interface())
				additionalTypes = append(additionalTypes, add...)

				line := fmt.Sprintf("%v: %v", key, res)
				lines = append(lines, line)
				continue
			}

			if valKind == reflect.Slice {
				add, res := StructToTYPES(val.Interface())
				additionalTypes = append(additionalTypes, add...)

				line := fmt.Sprintf("%v: %v", key, res)
				lines = append(lines, line)
				continue
			}

			if valKind == reflect.Ptr {
				nt := reflect.New(ty.Field(i).Type.Elem())
				nt = reflect.Indirect(nt)

				add, res := StructToTYPES(nt.Interface())
				additionalTypes = append(additionalTypes, add...)

				line := fmt.Sprintf("%v: %v", key, fmt.Sprintf("%v | undefined", res))
				lines = append(lines, line)
				continue
			}

			if valKind == reflect.Map {
				value := "unknown"

				if val.Type().Elem().Kind() == reflect.Ptr {
					if val.Type().Elem().Kind() == reflect.Ptr {
						val := val.Type().Elem().Elem()
						n := reflect.New(val)
						n = reflect.Indirect(n)

						add, res := StructToTYPES(n.Interface())
						additionalTypes = append(additionalTypes, add...)

						value = res
					} else {
						prim := primitives(val.Type().Elem(), val.Type().Elem().Kind(), true)

						if prim != "" {
							value = prim
						}
					}
				} else if val.Type().Elem().Kind() == reflect.Struct {
					val := reflect.New(val.Type().Elem())
					val = reflect.Indirect(val)

					add, res := StructToTYPES(val.Interface())
					additionalTypes = append(additionalTypes, add...)

					value = res
				} else {
					prim := primitives(val.Type(), val.Type().Elem().Kind(), true)
					if prim != "" {
						value = prim
					}
				}

				keyin := primitives(val.Type(), val.Type().Key().Kind(), false)

				line := fmt.Sprintf("{ [ key: %v ]: %v }", keyin, value)
				line = fmt.Sprintf("%v: %v", key, line)
				lines = append(lines, line)

				continue
			}

			line := fmt.Sprintf("%v: %v", key, "unknown")
			lines = append(lines, line)
		}

		return nil, fmt.Sprint("{ ", strings.Join(lines, "; "), " }")
	}

	return
}

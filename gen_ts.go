package goopy

import (
	"fmt"
	ref "reflect"
	"strings"

	"golang.org/x/exp/slices"
)

func gen_ts(t interface{}, indent int, constantIndent int) (res string) {
	res = "unknown"
	values := ref.ValueOf(t)
	kind := values.Kind()

	if kind == ref.Invalid {
		res = "null"
		return
	}

	if slices.Contains(primitiveKinds, kind) {
		if kind == ref.String {
			res = "string"
			return
		}

		if kind == ref.Bool {
			res = "boolean"
			return
		}

		if slices.Contains(numberKinds, kind) {
			res = "number"
			return
		}
	}

	if kind == ref.Ptr || kind == ref.Slice {
		elem := ref.TypeOf(t).Elem()
		value := ref.Indirect(ref.New(elem))
		types := gen_ts(value.Interface(), indent, constantIndent)

		if kind == ref.Ptr {
			res = fmt.Sprintf("%v | undefined", types)
		}

		if kind == ref.Slice {
			res = fmt.Sprintf("Array<%v>", types)
		}

		return
	}

	if kind == ref.Map {
		newType := ref.TypeOf(t)
		key := ref.Indirect(ref.New(newType.Key()))
		elem := ref.Indirect(ref.New(newType.Elem()))

		if key.Kind() == ref.String || slices.Contains(numberKinds, key.Kind()) {
			keyType := gen_ts(key.Interface(), indent, constantIndent)
			elemType := gen_ts(elem.Interface(), indent, constantIndent)

			res = fmt.Sprintf("Record<%v, %v>", keyType, elemType)
			return
		}
	}

	if kind == ref.Struct {
		keys := ref.TypeOf(t)
		indentString := strings.Repeat(" ", indent)

		eolIndent := indent - constantIndent
		if indent-constantIndent <= 0 {
			eolIndent = 0
		}

		eolString := strings.Repeat(" ", eolIndent)

		lines := []string{}
		for i := 0; i < values.NumField(); i++ {
			key := keys.Field(i).Name
			value := "unknown"
			if tag := keys.Field(i).Tag.Get("json"); tag != "" {
				key = tag
			}

			if tag := keys.Field(i).Tag.Get("alter"); tag != "" {
				value = tag
				line := fmt.Sprintf("%v: %v", key, value)
				lines = append(lines, indentString+line)
				continue
			}

			value = gen_ts(values.Field(i).Interface(), indent+constantIndent, constantIndent)
			if values.Field(i).Kind() == ref.Ptr {
				key = key + "?"
				value = strings.ReplaceAll(value, " | undefined", "")
			}

			if values.Field(i).Kind() == ref.Interface {
				elem := values.Field(i).Elem()
				if elem.Kind() == ref.Ptr {
					key = key + "?"
					value = strings.ReplaceAll(value, " | undefined", "")

				}
				// fmt.Println(key, value, elem)

			}

			if tag := keys.Field(i).Tag.Get("alter"); tag != "" {
				value = tag
			}

			line := fmt.Sprintf("%v: %v", key, value)
			lines = append(lines, indentString+line)
		}

		res = fmt.Sprintf("{\n%v\n%v}", strings.Join(lines, "\n"), eolString)
		return
	}

	return
}

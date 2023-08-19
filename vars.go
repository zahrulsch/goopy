package goopy

import "reflect"

var primitiveKinds = []reflect.Kind{
	reflect.String,
	reflect.Bool,
	reflect.Int,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
}

var numberKinds = []reflect.Kind{
	reflect.Int,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
}

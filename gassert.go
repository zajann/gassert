package gassert

import (
	"fmt"
	"reflect"
)

type assert struct {
	r   bool
	err error
}

var a assert

func Zeros(xs ...interface{}) *assert {
	a.zeros(xs...)
	return &a
}

func Equals(x interface{}, y interface{}) *assert {
	a.equals(x, y)
	return &a
}

func Less(x interface{}, y interface{}) *assert {
	return nil
}

func Greater(x interface{}, y interface{}) *assert {
	return nil
}

func (a *assert) Go() error {
	defer a.clear()
	if a.r {
		return a.err
	}
	return nil
}

func (a *assert) clear() {
	a.r = true
	a.err = nil
}

func (a *assert) zeros(xs ...interface{}) {
	for _, x := range xs {
		val := reflect.ValueOf(x)
		if val.IsZero() {
			a.r = true
			a.err = fmt.Errorf("zero-values in %s", val.Kind())
		}
	}
}

func (a *assert) equals(x interface{}, y interface{}) {
	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)

	if xVal.Kind() != yVal.Kind() {
		panic("gassert.Equals: cannot compare, different types")
	}

	switch xVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// compare Int

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// compoare Uint
	case reflect.Float32, reflect.Float64:
		// compare Float
	case reflect.String:
		// compare string
	case reflect.Array:
		// compare array
	case reflect.Slice:
		// compare slice
	case reflect.Map:
		// compare map
	case reflect.Ptr:
		// compoare pointer
	case reflect.Struct:
		// compare struct
	default:
		panic(fmt.Sprintf("gassert.Equals: unsupported type [%s]", xVal.Kind().String()))
	}

}

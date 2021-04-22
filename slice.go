package gassert

import "reflect"

func isSlice(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.Slice
}

func sliceCapEquals(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Cap() == b
}

func sliceCapLess(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Cap() < b
}

func sliceCapLessOrEquals(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Cap() <= b
}

func sliceCapGreater(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Cap() > b
}

func sliceCapGreaterOrEquals(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Cap() >= b
}

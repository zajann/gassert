package gassert

import "reflect"

func lenEquals(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Len() == b
}

func lenLess(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Len() < b
}

func lenLessOrEquals(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Len() <= b
}

func lenGreater(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Len() > b
}

func lenGreaterOrEquals(a interface{}, b int) bool {
	val := reflect.ValueOf(a)
	return val.Len() >= b
}

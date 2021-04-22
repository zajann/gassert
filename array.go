package gassert

import "reflect"

func isArray(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.Array
}

package gassert

import "reflect"

func isMap(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.Map
}

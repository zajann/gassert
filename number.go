package gassert

import (
	"reflect"
)

var numberMap map[reflect.Kind]struct{}

func init() {
	numberMap = make(map[reflect.Kind]struct{})
	numberMap[reflect.Int] = struct{}{}
	numberMap[reflect.Int8] = struct{}{}
	numberMap[reflect.Int16] = struct{}{}
	numberMap[reflect.Int32] = struct{}{}
	numberMap[reflect.Int64] = struct{}{}
	numberMap[reflect.Uint] = struct{}{}
	numberMap[reflect.Uint8] = struct{}{}
	numberMap[reflect.Uint16] = struct{}{}
	numberMap[reflect.Uint32] = struct{}{}
	numberMap[reflect.Uint64] = struct{}{}
	numberMap[reflect.Float32] = struct{}{}
	numberMap[reflect.Float64] = struct{}{}
}

func isNumber(x reflect.Value) bool {
	_, ok := numberMap[x.Kind()]
	return ok
}

func parseToFloat64s(x, y interface{}) (float64, float64, bool) {
	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)

	if isNumber(xVal) && isNumber(yVal) {
		return parseToFloat64FromNumber(xVal), parseToFloat64FromNumber(yVal), true
	}

	return 0, 0, false
}

func parseToFloat64FromNumber(x reflect.Value) float64 {
	var f float64

	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f = float64(x.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f = float64(x.Uint())
	case reflect.Float32, reflect.Float64:
		f = x.Float()
	}

	return f
}

func float64Equals(a float64, b float64) bool {
	return a == b
}

func float64Zero(a float64) bool {
	return float64Equals(a, 0)
}

func float64Greater(a float64, b float64) bool {
	return a > b
}

func float64GreaterOrEquals(a float64, b float64) bool {
	return float64Equals(a, b) || float64Greater(a, b)
}

func float64Less(a float64, b float64) bool {
	return a < b
}

func float64LessOrEquals(a float64, b float64) bool {
	return float64Equals(a, b) || float64Less(a, b)
}

func uint64Equals(a uint64, b uint64) bool {
	return a == b
}

func uint64Zero(a uint64) bool {
	return uint64Equals(a, 0)
}

func uint64Greater(a uint64, b uint64) bool {
	return a > b
}

func uint64GreaterOrEquals(a uint64, b uint64) bool {
	return uint64Equals(a, b) || uint64Greater(a, b)
}

func uint64Less(a uint64, b uint64) bool {
	return a < b
}

func uint64LessOrEquals(a uint64, b uint64) bool {
	return uint64Equals(a, b) || uint64Less(a, b)
}

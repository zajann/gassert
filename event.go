package gassert

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

// panic only here

var eventPool = &sync.Pool{
	New: func() interface{} {
		return &Event{
			errs: make([]error, 0, 5),
		}
	},
}

func newEvent() *Event {
	e := eventPool.Get().(*Event)
	e.errs = e.errs[:0]
	return e
}

func putEvent(e *Event) {
	if e == nil {
		return
	}
	eventPool.Put(e)
}

type Event struct {
	errs []error
}

func (e *Event) Zeros(xs ...interface{}) *Event {
	for _, x := range xs {
		val := reflect.ValueOf(x)
		if val.IsZero() {
			e.addError("zero-values in %s", val.Kind().String())
		}
	}
	return e
}

func (e *Event) DeepEqual(x, y interface{}) *Event {
	if reflect.DeepEqual(x, y) {
		e.addError("deep equals")
	}
	return e
}

func (e *Event) Equals(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64Equals(xx, yy) {
			e.addError("number less")
		}
	} else {
		xVal := reflect.ValueOf(x)
		switch xVal.Kind() {
		case reflect.String:
			if yy, ok := y.(string); ok {
				if stringEquals(xVal.String(), yy) {
					e.addError("string equals")
				}
			}
		default:
			e.DeepEqual(x, y)
		}
	}
	return e
}

func (e *Event) NotEquals(x, y interface{}) *Event {
	before := len(e.errs)
	e = e.Equals(x, y)
	after := len(e.errs)

	if after == before {
		e.addError("not equals")
	} else {
		e.errs = e.errs[:after-1]
	}
	return e
}

func (e *Event) NumLess(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64Less(xx, yy) {
			e.addError("number less")
		}
	} else {
		panic("gassert.Less: only numbers available")
	}
	return e
}

func (e *Event) NumLessOrEquals(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64LessOrEquals(xx, yy) {
			e.addError("number less")
		}
	} else {
		panic("gassert.LessOrEquals: only numbers available")
	}
	return e
}

func (e *Event) NumGreater(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64Greater(xx, yy) {
			e.addError("number less")
		}
	} else {
		panic("gassert.Greater: only numbers available")
	}
	return e
}

func (e *Event) NumGreaterOrEquals(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64GreaterOrEquals(xx, yy) {
			e.addError("number less")
		}
	} else {
		panic("gassert.GreaterOrEquals: only numbers acceptable")
	}
	return e
}

func (e *Event) StrLenEquals(s string, n int) *Event {
	if lenEquals(s, n) {
		e.addError("string length equals")
	}
	return e
}

func (e *Event) StrLenNotEquals(s string, n int) *Event {
	if !lenEquals(s, n) {
		e.addError("string length not equals")
	}
	return e
}

func (e *Event) StrLenLess(s string, n int) *Event {
	if lenLess(s, n) {
		e.addError(" length less than")
	}
	return e
}

func (e *Event) StrLenLessOrEquals(s string, n int) *Event {
	if lenLessOrEquals(s, n) {
		e.addError(" length less than")
	}
	return e
}

func (e *Event) StrLenGreater(s string, n int) *Event {
	if lenGreater(s, n) {
		e.addError(" length less than")
	}
	return e
}

func (e *Event) StrLenGreaterOrEquals(s string, n int) *Event {
	if lenGreaterOrEquals(s, n) {
		e.addError(" length less than")
	}
	return e
}

func (e *Event) ArrLenEquals(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.ArrLenEquals: only array type acceptable")
	}

	if lenEquals(x, n) {
		e.addError("array length equals")
	}
	return e
}

func (e *Event) ArrLenNotEquals(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.ArrLenNotEquals: only array type acceptable")
	}

	if !lenEquals(x, n) {
		e.addError("array length not equals")
	}
	return e
}

func (e *Event) ArrLenLess(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.ArrLenLess: only array type acceptable")
	}

	if lenLess(x, n) {
		e.addError("array length equals")
	}
	return e
}

func (e *Event) ArrLenLessOrEquals(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.ArrLenLessOrEquals: only array type acceptable")
	}

	if lenLessOrEquals(x, n) {
		e.addError("array length equals")
	}
	return e
}

func (e *Event) ArrLenGreater(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.ArrLenGreater: only array type acceptable")
	}

	if lenGreater(x, n) {
		e.addError("array length equals")
	}
	return e
}

func (e *Event) ArrLenGreaterOrEquals(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.ArrLenGreaterOrEquals: only array type acceptable")
	}

	if lenGreaterOrEquals(x, n) {
		e.addError("array length equals")
	}
	return e
}

func (e *Event) SliceLenEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceLenEquals: only slice type acceptable")
	}

	if lenEquals(x, n) {
		e.addError("slice length equals")
	}
	return e
}

func (e *Event) SliceLenNotEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceLenNotEquals: only slice type acceptable")
	}

	if !lenEquals(x, n) {
		e.addError("slice length not equals")
	}
	return e
}

func (e *Event) SliceLenLess(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceLenLess: only slice type acceptable")
	}

	if lenLess(x, n) {
		e.addError("slice length equals")
	}
	return e
}

func (e *Event) SliceLenLessOrEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceLenLessOrEquals: only slice type acceptable")
	}

	if lenLessOrEquals(x, n) {
		e.addError("slice length equals")
	}
	return e
}

func (e *Event) SliceLenGreater(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceLenGreater: only slice type acceptable")
	}

	if lenGreater(x, n) {
		e.addError("slice length equals")
	}
	return e
}

func (e *Event) SliceLenGreaterOrEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceLenGreaterOrEquals: only slice type acceptable")
	}

	if lenGreaterOrEquals(x, n) {
		e.addError("slice length equals")
	}
	return e
}

func (e *Event) SliceCapEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceCapEquals: only slice type acceptable")
	}

	if sliceCapEquals(x, n) {
		e.addError("slice capacity equals")
	}
	return e
}

func (e *Event) SliceCapNotEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceCapNotEquals: only slice type acceptable")
	}

	if !sliceCapEquals(x, n) {
		e.addError("slice capacity not equals")
	}
	return e
}

func (e *Event) SliceCapLess(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceCapLess: only slice type acceptable")
	}

	if sliceCapLess(x, n) {
		e.addError("slice capacity equals")
	}
	return e
}

func (e *Event) SliceCapLessOrEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceCapLessOrEquals: only slice type acceptable")
	}

	if sliceCapLessOrEquals(x, n) {
		e.addError("slice capacity equals")
	}
	return e
}

func (e *Event) SliceCapGreater(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceCapGreater: only slice type acceptable")
	}

	if sliceCapGreater(x, n) {
		e.addError("slice capacity equals")
	}
	return e
}

func (e *Event) SliceCapGreaterOrEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.SliceCapGreaterOrEquals: only slice type acceptable")
	}

	if sliceCapGreaterOrEquals(x, n) {
		e.addError("slice capacity equals")
	}
	return e
}

func (e *Event) MapLenEquals(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.MapLenEquals: only map type acceptable")
	}

	if lenEquals(x, n) {
		e.addError("map length equals")
	}
	return e
}

func (e *Event) MapLenNotEquals(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.MapLenNotEquals: only map type acceptable")
	}

	if !lenEquals(x, n) {
		e.addError("map length not equals")
	}
	return e
}

func (e *Event) MapLenLess(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.MapLenLess: only map type acceptable")
	}

	if lenLess(x, n) {
		e.addError("map length equals")
	}
	return e
}

func (e *Event) MapLenLessOrEquals(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.MapLenLessOrEquals: only map type acceptable")
	}

	if lenLessOrEquals(x, n) {
		e.addError("map length equals")
	}
	return e
}

func (e *Event) MapLenGreater(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.MapLenGreater: only map type acceptable")
	}

	if lenGreater(x, n) {
		e.addError("map length equals")
	}
	return e
}

func (e *Event) MapLenGreaterOrEquals(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.MapLenGreaterOrEquals: only map type acceptable")
	}

	if lenGreaterOrEquals(x, n) {
		e.addError("map length equals")
	}
	return e
}

func (e *Event) Panic() {
	defer putEvent(e)
	if len(e.errs) > 0 {
		panic("gAssert trigger panic\n\tdfdf\n")
	}
}

func (e *Event) Err() error {
	defer putEvent(e)
	if len(e.errs) > 0 {
		pc, _, _, _ := runtime.Caller(1)
		callerName := runtime.FuncForPC(pc).Name()

		return fmt.Errorf("gAssertError: %s", callerName)
	}
	return nil
}

func (e *Event) addError(format string, a ...interface{}) {
	e.errs = append(e.errs, fmt.Errorf(format, a...))
}

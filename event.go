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
			e.appendError(zeroValueError(val.Kind().String()))
		}
	}
	return e
}

func (e *Event) DeepEqual(x, y interface{}) *Event {
	if reflect.DeepEqual(x, y) {
		e.appendError(equalsError(reflect.ValueOf(x).Kind().String()))
	}
	return e
}

func (e *Event) Equals(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64Equals(xx, yy) {
			e.appendError(equalsError("number"))
		}
	} else {
		xVal := reflect.ValueOf(x)
		switch xVal.Kind() {
		case reflect.String:
			if yy, ok := y.(string); ok {
				if stringEquals(xVal.String(), yy) {
					e.appendError(equalsError("string"))
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
		val := reflect.ValueOf(x)
		kind := val.Kind().String()
		if isNumber(val) {
			kind = "number"
		}
		e.appendError(notEqualsError(kind))
	} else {
		e.errs = e.errs[:after-1]
	}
	return e
}

func (e *Event) NumLess(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64Less(xx, yy) {
			e.appendError(numLessError)
		}
	} else {
		panic("gassert.Event.Less: only numbers acceptable")
	}
	return e
}

func (e *Event) NumLessOrEquals(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64LessOrEquals(xx, yy) {
			e.appendError(numLessOrEqualsError)
		}
	} else {
		panic("gassert.Event.LessOrEquals: only numbers acceptable")
	}
	return e
}

func (e *Event) NumGreater(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64Greater(xx, yy) {
			e.appendError(numGreaterError)
		}
	} else {
		panic("gassert.Event.Greater: only numbers acceptable")
	}
	return e
}

func (e *Event) NumGreaterOrEquals(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64GreaterOrEquals(xx, yy) {
			e.appendError(numGreaterOrEqualsError)
		}
	} else {
		panic("gassert.Event.GreaterOrEquals: only numbers acceptable")
	}
	return e
}

func (e *Event) StrLenEquals(s string, n int) *Event {
	if lenEquals(s, n) {
		e.appendError(lenEqualsError("string"))
	}
	return e
}

func (e *Event) StrLenNotEquals(s string, n int) *Event {
	if !lenEquals(s, n) {
		e.appendError(lenNotEqualsError("string"))
	}
	return e
}

func (e *Event) StrLenLess(s string, n int) *Event {
	if lenLess(s, n) {
		e.appendError(lenLessError("string"))
	}
	return e
}

func (e *Event) StrLenLessOrEquals(s string, n int) *Event {
	if lenLessOrEquals(s, n) {
		e.appendError(lenLessOrEqualsError("string"))
	}
	return e
}

func (e *Event) StrLenGreater(s string, n int) *Event {
	if lenGreater(s, n) {
		e.appendError(lenGreaterError("string"))
	}
	return e
}

func (e *Event) StrLenGreaterOrEquals(s string, n int) *Event {
	if lenGreaterOrEquals(s, n) {
		e.appendError(lenGreaterOrEqualsError("string"))
	}
	return e
}

func (e *Event) ArrLenEquals(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.Event.ArrLenEquals: only array type acceptable")
	}

	if lenEquals(x, n) {
		e.appendError(lenEqualsError("array"))
	}
	return e
}

func (e *Event) ArrLenNotEquals(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.Event.ArrLenNotEquals: only array type acceptable")
	}

	if !lenEquals(x, n) {
		e.appendError(lenNotEqualsError("array"))
	}
	return e
}

func (e *Event) ArrLenLess(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.Event.ArrLenLess: only array type acceptable")
	}

	if lenLess(x, n) {
		e.appendError(lenLessError("array"))
	}
	return e
}

func (e *Event) ArrLenLessOrEquals(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.Event.ArrLenLessOrEquals: only array type acceptable")
	}

	if lenLessOrEquals(x, n) {
		e.appendError(lenLessOrEqualsError("array"))
	}
	return e
}

func (e *Event) ArrLenGreater(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.Event.ArrLenGreater: only array type acceptable")
	}

	if lenGreater(x, n) {
		e.appendError(lenGreaterError("array"))
	}
	return e
}

func (e *Event) ArrLenGreaterOrEquals(x interface{}, n int) *Event {
	if !isArray(x) {
		panic("gassert.Event.ArrLenGreaterOrEquals: only array type acceptable")
	}

	if lenGreaterOrEquals(x, n) {
		e.appendError(lenGreaterOrEqualsError("array"))
	}
	return e
}

func (e *Event) SliceLenEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceLenEquals: only slice type acceptable")
	}

	if lenEquals(x, n) {
		e.appendError(lenEqualsError("slice"))
	}
	return e
}

func (e *Event) SliceLenNotEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceLenNotEquals: only slice type acceptable")
	}

	if !lenEquals(x, n) {
		e.appendError(lenNotEqualsError("slice"))
	}
	return e
}

func (e *Event) SliceLenLess(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceLenLess: only slice type acceptable")
	}

	if lenLess(x, n) {
		e.appendError(lenLessError("slice"))
	}
	return e
}

func (e *Event) SliceLenLessOrEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceLenLessOrEquals: only slice type acceptable")
	}

	if lenLessOrEquals(x, n) {
		e.appendError(lenLessOrEqualsError("slice"))
	}
	return e
}

func (e *Event) SliceLenGreater(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceLenGreater: only slice type acceptable")
	}

	if lenGreater(x, n) {
		e.appendError(lenGreaterError("slice"))
	}
	return e
}

func (e *Event) SliceLenGreaterOrEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceLenGreaterOrEquals: only slice type acceptable")
	}

	if lenGreaterOrEquals(x, n) {
		e.appendError(lenGreaterOrEqualsError("slice"))
	}
	return e
}

func (e *Event) SliceCapEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceCapEquals: only slice type acceptable")
	}

	if sliceCapEquals(x, n) {
		e.appendError(capEqualsError("slice"))
	}
	return e
}

func (e *Event) SliceCapNotEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceCapNotEquals: only slice type acceptable")
	}

	if !sliceCapEquals(x, n) {
		e.appendError(capNotEqualsError("slice"))
	}
	return e
}

func (e *Event) SliceCapLess(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceCapLess: only slice type acceptable")
	}

	if sliceCapLess(x, n) {
		e.appendError(capLessError("slice"))
	}
	return e
}

func (e *Event) SliceCapLessOrEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceCapLessOrEquals: only slice type acceptable")
	}

	if sliceCapLessOrEquals(x, n) {
		e.appendError(capLessOrEqualsError("slice"))
	}
	return e
}

func (e *Event) SliceCapGreater(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceCapGreater: only slice type acceptable")
	}

	if sliceCapGreater(x, n) {
		e.appendError(capGreaterError("slice"))
	}
	return e
}

func (e *Event) SliceCapGreaterOrEquals(x interface{}, n int) *Event {
	if !isSlice(x) {
		panic("gassert.Event.SliceCapGreaterOrEquals: only slice type acceptable")
	}

	if sliceCapGreaterOrEquals(x, n) {
		e.appendError(capGreaterOrEqualsError("slice"))
	}
	return e
}

func (e *Event) MapLenEquals(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.Event.MapLenEquals: only map type acceptable")
	}

	if lenEquals(x, n) {
		e.appendError(lenEqualsError("map"))
	}
	return e
}

func (e *Event) MapLenNotEquals(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.Event.MapLenNotEquals: only map type acceptable")
	}

	if !lenEquals(x, n) {
		e.appendError(lenNotEqualsError("map"))
	}
	return e
}

func (e *Event) MapLenLess(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.Event.MapLenLess: only map type acceptable")
	}

	if lenLess(x, n) {
		e.appendError(lenLessError("map"))
	}
	return e
}

func (e *Event) MapLenLessOrEquals(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.Event.MapLenLessOrEquals: only map type acceptable")
	}

	if lenLessOrEquals(x, n) {
		e.appendError(lenLessOrEqualsError("map"))
	}
	return e
}

func (e *Event) MapLenGreater(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.Event.MapLenGreater: only map type acceptable")
	}

	if lenGreater(x, n) {
		e.appendError(lenGreaterError("map"))
	}
	return e
}

func (e *Event) MapLenGreaterOrEquals(x interface{}, n int) *Event {
	if !isMap(x) {
		panic("gassert.Event.MapLenGreaterOrEquals: only map type acceptable")
	}

	if lenGreaterOrEquals(x, n) {
		e.appendError(lenGreaterOrEqualsError("map"))
	}
	return e
}

func (e *Event) Panic() {
	defer putEvent(e)
	if len(e.errs) > 0 {
		panic("gAssertError")
	}
}

func (e *Event) PanicDetails() {
	defer putEvent(e)
	if len(e.errs) > 0 {
		panic(Errors(e.errs))
	}

}

func (e *Event) Err() error {
	defer putEvent(e)
	if len(e.errs) > 0 {
		pc, _, line, _ := runtime.Caller(1)
		callerName := runtime.FuncForPC(pc).Name()
		return fmt.Errorf("gAssertError: %s, line %d", callerName, line)
	}
	return nil
}

func (e *Event) ErrDetails() error {
	defer putEvent(e)
	if len(e.errs) > 0 {
		return Errors(e.errs)
	}
	return nil
}

func (e *Event) appendError(err error) {
	e.errs = append(e.errs, err)
}

package gassert

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

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
				if xVal.String() == yy {
					e.addError("string equals")
				}
			}
		default:
			e.DeepEqual(x, y)
		}
	}
	return e
}

func (e *Event) Less(x, y interface{}) *Event {
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

func (e *Event) LessOrEquals(x, y interface{}) *Event {
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

func (e *Event) Greater(x, y interface{}) *Event {
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

func (e *Event) GreaterOrEquals(x, y interface{}) *Event {
	xx, yy, ok := parseToFloat64s(x, y)
	if ok {
		if float64GreaterOrEquals(xx, yy) {
			e.addError("number less")
		}
	} else {
		panic("gassert.GreaterOrEquals: only numbers available")
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
		return errors.New("error")
	}
	return nil
}

func (e *Event) addError(format string, a ...interface{}) {
	e.errs = append(e.errs, fmt.Errorf(format, a...))
}

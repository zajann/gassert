package gassert

func Go(condition bool) {
	if condition {
		panic("gAssertError")
	}
}

func Zeros(xs ...interface{}) {
	New().Zeros(xs...).Panic()
}

func Equals(x, y interface{}) {
	New().Equals(x, y).Panic()
}

func NotEquals(x, y interface{}) {
	New().NotEquals(x, y).Panic()
}

func New() *Event {
	return newEvent()
}

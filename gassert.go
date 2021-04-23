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

func NumLess(x, y interface{}) {
	New().NumLess(x, y).Panic()
}

func NumLessOrEquals(x, y interface{}) {
	New().NumLessOrEquals(x, y).Panic()
}

func NumGreater(x, y interface{}) {
	New().NumGreater(x, y).Panic()

}

func NumGreaterOrEquals(x, y interface{}) {
	New().NumGreaterOrEquals(x, y).Panic()
}

func StrLenEquals(s string, n int) {
	New().StrLenEquals(s, n).Panic()
}

func StrLenNotEquals(s string, n int) {
	New().StrLenNotEquals(s, n).Panic()
}

func StrLenLess(s string, n int) {
	New().StrLenLess(s, n).Panic()
}

func StrLenLessOrEquals(s string, n int) {
	New().StrLenLessOrEquals(s, n).Panic()
}

func StrLenGreater(s string, n int) {
	New().StrLenGreater(s, n).Panic()
}

func StrLenGreaterOrEquals(s string, n int) {
	New().StrLenGreaterOrEquals(s, n).Panic()
}

func ArrLenEquals(s interface{}, n int) {
	New().ArrLenEquals(s, n).Panic()
}

func ArrLenNotEquals(s interface{}, n int) {
	New().ArrLenNotEquals(s, n).Panic()
}

func ArrLenLess(s interface{}, n int) {
	New().ArrLenLess(s, n).Panic()
}

func ArrLenLessOrEquals(s interface{}, n int) {
	New().ArrLenLessOrEquals(s, n).Panic()
}

func ArrLenGreater(s interface{}, n int) {
	New().ArrLenGreater(s, n).Panic()
}

func ArrLenGreaterOrEquals(s interface{}, n int) {
	New().ArrLenGreaterOrEquals(s, n).Panic()
}

func SliceLenEquals(s interface{}, n int) {
	New().SliceLenEquals(s, n).Panic()
}

func SliceLenNotEquals(s interface{}, n int) {
	New().SliceLenNotEquals(s, n).Panic()
}

func SliceLenLess(s interface{}, n int) {
	New().SliceLenLess(s, n).Panic()
}

func SliceLenLessOrEquals(s interface{}, n int) {
	New().SliceLenLessOrEquals(s, n).Panic()
}

func SliceCapLess(s interface{}, n int) {
	New().SliceCapLess(s, n).Panic()
}

func SliceCapLessOrEquals(s interface{}, n int) {
	New().SliceCapLessOrEquals(s, n).Panic()
}

func SliceCapGreater(s interface{}, n int) {
	New().SliceCapGreater(s, n).Panic()
}

func SliceCapGreaterOrEquals(s interface{}, n int) {
	New().SliceCapGreaterOrEquals(s, n).Panic()
}

func MapLenEquals(s interface{}, n int) {
	New().MapLenEquals(s, n).Panic()
}

func MapLenNotEquals(s interface{}, n int) {
	New().MapLenNotEquals(s, n).Panic()
}

func MapLenLess(s interface{}, n int) {
	New().MapLenLess(s, n).Panic()
}

func MapLenLessOrEquals(s interface{}, n int) {
	New().MapLenLessOrEquals(s, n).Panic()
}

func MapLenGreater(s interface{}, n int) {
	New().MapLenGreater(s, n).Panic()
}

func MapLenGreaterOrEquals(s interface{}, n int) {
	New().MapLenGreaterOrEquals(s, n).Panic()
}

func New() *Event {
	return newEvent()
}

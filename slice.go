package gassert

// FIXME: ALL
func sliceLenEquals(x []interface{}, n int) bool {
	return len(x) == n
}

func sliceLenLess(x []interface{}, n int) bool {
	return len(x) < n
}

func sliceLenLessOrEquals(x []interface{}, n int) bool {
	return len(x) <= n
}

func sliceLenGreater(x []interface{}, n int) bool {
	return len(x) > n
}

func sliceLenGreaterOrEquals(x []interface{}, n int) bool {
	return len(x) >= n
}

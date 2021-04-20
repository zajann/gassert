package gassert

func stringEquals(a, b string) bool {
	return a == b
}

func stringLenEquals(a, b string) bool {
	return len(a) == len(b)
}

func stringLenLess(a string, b int) bool {
	return len(a) < b
}

func stringLenLessOrEquals(a string, b int) bool {
	return len(a) <= b
}

func stringLenGreater(a string, b int) bool {
	return len(a) > b
}

func stringLenGreaterOrEquals(a string, b int) bool {
	return len(a) >= b
}

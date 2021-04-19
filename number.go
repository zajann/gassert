package gassert

func intEquals(a int, b int) bool {
	return a == b
}

func intZero(a int) bool {
	return intEquals(a, 0)
}

func intGreater(a int, b int) bool {
	return a > b
}

func intGreaterOrEquals(a int, b int) bool {
	return intEquals(a, b) || intGreater(a, b)
}

func intLess(a int, b int) bool {
	return a < b
}

func intLessOrEquals(a int, b int) bool {
	return intEquals(a, b) || intLess(a, b)
}

func uintEquals(a uint, b uint) bool {
	return a == b
}

func uintZero(a uint) bool {
	return uintEquals(a, 0)
}

func uintGreater(a uint, b uint) bool {
	return a > b
}

func uintGreaterOrEquals(a uint, b uint) bool {
	return uintEquals(a, b) || uintGreater(a, b)
}

func uintLess(a uint, b uint) bool {
	return a < b
}

func uintLessOrEquals(a uint, b uint) bool {
	return uintEquals(a, b) || uintLess(a, b)
}

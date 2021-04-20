package main

import (
	"fmt"

	"github.com/zajann/gassert"
)

func main() {

	i := 99

	TestFunc(i)

}

func TestFunc(n int) {
	var t int32
	t = 40

	if err := gassert.New().NumGreater(n, t).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().NumLess(n, 100).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().NumLessOrEquals(n, 99).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().NumLessOrEquals(n, 99.0).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().NumGreater(n, 50).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().NumGreaterOrEquals(n, 99).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().NumGreaterOrEquals(n, 99.0).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().Equals(n, 99).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().Equals(n, 99.0).Err(); err != nil {
		fmt.Println(err)
	}

}

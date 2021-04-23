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
	// var t int32
	// t = 40

	if err := gassert.New().NotEquals(n, 100).Err(); err != nil {
		fmt.Println(err)
	}

	gassert.New().NumGreater(n, 89.9999).Panic()
}

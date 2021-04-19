package main

import (
	"fmt"

	"github.com/zajann/gassert"
)

func main() {
	var test uint8

	if err := gassert.Zeros(test).Go(); err != nil {
		panic(err)
	}

	fmt.Println(test)
}

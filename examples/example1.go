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
	if err := gassert.New().Less(n, 100).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().LessOrEquals(n, 99).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().LessOrEquals(n, 99.0).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().Greater(n, 50).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().GreaterOrEquals(n, 99).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().GreaterOrEquals(n, 99.0).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().Equals(n, 99).Err(); err != nil {
		fmt.Println(err)
	}
	if err := gassert.New().Equals(n, 99.0).Err(); err != nil {
		fmt.Println(err)
	}

}

package main

import (
	"log"

	"github.com/zajann/gassert"
)

func main() {
	n := 15
	s := "abc"
	gassert.Go((n-5)%10 == 0 || len(s)*2 > 10)

}

func SingleWithPanic(n int) {
	gassert.Go(n > 1)

	gassert.Equals(n, 10)

	gassert.NotEquals(n, 11)

	gassert.Zeros(n)
}

func MultiWithError(n int, s string, ss []string, m map[string]int) {
	gassert.New().Zeros(ss, m).NumLess(n, 0).StrLenEquals(s, 5).Panic()

	if err := gassert.New().Zeros(s, m).NumLess(n, 0).SliceLenLess(ss, 10).Err(); err != nil {
		log.Println(err)
	}
}

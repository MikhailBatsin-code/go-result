package main

import (
	"fmt"

	"github.com/MikhailBatsin-code/go-result"
)

func definitelyOk() result.Result[int] {
	return result.Ok(10)
}

func main() {
	res := definitelyOk()

	// get ok value
	// fmt.Println(res.Unwrap())

	// conditions
	if res.IsOk() {
		fmt.Println(res.Unwrap())
	}

	if res.IsOkAnd(func(i int) bool { return i%2 == 0 }) {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

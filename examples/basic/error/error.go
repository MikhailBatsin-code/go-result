package main

import (
	"errors"

	"github.com/MikhailBatsin-code/go-result"
)

func definitelyError() result.Result[int] {
	return result.Error[int](errors.New("its an error"))
}

func main() {
	res := definitelyError()

	// will exit
	res.Unwrap()

	// will just print error
	// res.UnwrapSoft()

	// will call callback
	// res.UnwrapOr(func(err error) {
	// 	fmt.Println("there is an error((")
	// })

	// will return 5 because of error
	// fmt.Println(res.UnwrapOrDefault(5))

	// will return error
	// fmt.Println(res.UnwrapError())

	// will print your message and exit program if there is an error or return ok value
	// res.Expect("there is an error and that's very sad")

	// will just print your message or return ok value
	// res.ExpectSoft("sad")

	// some conditions
	// if res.IsOk() {
	// 	fmt.Println(res.Unwrap())
	// } else {
	// 	fmt.Println(res.UnwrapError())
	// }
}

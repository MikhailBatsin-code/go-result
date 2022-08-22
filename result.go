package result

import "log"

// shortcut type for result with only errors
type None struct {
}

type Result[T any] struct {
	ok  T
	err error
}

// will exit if err not nil or return ok value
func (r *Result[T]) Unwrap() T {
	if r.err != nil {
		log.Fatalln(r.err)
	}

	return r.ok
}

// will print error and return default value for T if err not nil or return ok value
func (r *Result[T]) UnwrapSoft() T {
	if r.err != nil {
		log.Println(r.err)
		return *new(T)
	}

	return r.ok
}

// if there is an error then call callback and return default value for T
func (r *Result[T]) UnwrapOr(fn func(error)) T {
	if r.err != nil {
		fn(r.err)
		return *new(T)
	}

	return r.ok
}

// if there is an error return default value
func (r *Result[T]) UnwrapOrDefault(defaultValue T) T {
	if r.err != nil {
		return defaultValue
	}

	return r.ok
}

// returns error
func (r *Result[T]) UnwrapError() error {
	return r.err
}

// exit program with given message
func (r *Result[T]) Expect(msg string) T {
	if r.err != nil {
		log.Fatalln(msg)
	}

	return r.ok
}

// just print message if there is an error and return default value for T
func (r *Result[T]) ExpectSoft(msg string) T {
	if r.err != nil {
		log.Println(msg)
		return *new(T)
	}

	return r.ok
}

// checks if there is no error
func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

// checks if there is no error and your own check is ok
func (r *Result[T]) IsOkAnd(fn func(T) bool) bool {
	return r.IsOk() && fn(r.ok)
}

// create result with ok value and nil error
func Ok[T any](val T) Result[T] {
	return Result[T]{
		ok:  val,
		err: nil,
	}
}

// create result with error and garbage ok
func Error[T any](err error) Result[T] {
	return Result[T]{
		ok:  *new(T),
		err: err,
	}
}

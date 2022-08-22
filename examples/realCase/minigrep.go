package main

import "log"

func main() {
	log.SetFlags(0)
	log.SetPrefix("Error: ")

	res := newConfig()
	config := res.Unwrap()

	run(config)
}

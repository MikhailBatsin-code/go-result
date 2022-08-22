package main

import (
	"errors"
	"os"

	"github.com/MikhailBatsin-code/go-result"
)

type Config struct {
	filename string
	query    string
}

func newConfig() result.Result[Config] {
	config := Config{}

	if len(os.Args) < 3 {
		return result.Error[Config](errors.New("not enough arguments"))
	}

	config.filename = os.Args[1]
	config.query = os.Args[2]

	return result.Ok(config)
}

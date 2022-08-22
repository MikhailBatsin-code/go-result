package main

import (
	"io/ioutil"
	"strings"

	"github.com/MikhailBatsin-code/go-result"
)

func getFileBytes(filename string) result.Result[[]byte] {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return result.Error[[]byte](err)
	}

	return result.Ok(fileBytes)
}

func run(config Config) {
	res := getFileBytes(config.filename)
	contents := string(res.Unwrap())

	search(config, strings.Split(
		strings.ReplaceAll(contents, "\r\n", "\n"),
		"\n",
	))
}

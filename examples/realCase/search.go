package main

import (
	"fmt"
	"strings"
)

func search(config Config, lines []string) {
	for index, line := range lines {
		if strings.Contains(line, config.query) {
			fmt.Printf("line %d: %s\n", index, line)
		}
	}
}

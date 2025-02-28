package main

import (
	"strings"
)

func Repeat(character string, toRepeat int) string {
	var repeated strings.Builder
	for i := 0; i < toRepeat; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func main() {
}

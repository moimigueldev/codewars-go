package main

import (
	"strings"
)

func main() {
	solve("your code rocks")
}

func solve(s string) (out string) {
	in := strings.Split(s, " ")
	for i, word := range in {

		in[i] = reverse(word)

	}

	out = strings.Join(in[:], " ")

	return out
}

func reverse(s string) (out string) {
	for _, char := range s {
		out = string(char) + out
	}
	return out
}

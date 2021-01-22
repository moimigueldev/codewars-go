package main

import (
	"strings"
)

// This time no story, no theory. The examples below show you how to write function accum:

// Examples:

// accum("abcd") -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") -> "C-Ww-Aaa-Tttt"
// The parameter of accum is a string which includes only letters from a..z and A..Z.

func main() {
	solve("Abcd")
}

func solve(s string) string {
	arr := strings.Split(strings.ToLower(s), "")
	for i, ch := range arr {
		str := strings.Repeat(string(ch), i+1)
		arr[i] = strings.Title(str)
	}
	return strings.Join(arr, "-")
}

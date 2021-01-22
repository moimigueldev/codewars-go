package main

// In this Kata, you will be given an array of strings and your task is to remove all consecutive duplicate letters from each string in the array.

// For example:

// dup(["abracadabra","allottee","assessee"]) = ["abracadabra","alote","asese"].

// dup(["kelless","keenness"]) = ["keles","kenes"].

// Strings will be lowercase only, no spaces. See test cases for more examples.

// Good luck!

// If you like this Kata, please try:

func main() {
	dup([]string{"kelless", "keenness"})
	dup([]string{"abracadabra", "allottee", "assessee"})
}

func dup(list []string) []string {
	errCounter := false
	for i := 0; i < len(list); i++ {
		for index, _ := range list[i] {
			if index+1 < len(list[i]) {
				current := string(list[i][index])
				next := string(list[i][index+1])
				if current == next {
					list[i] = list[i][:index+1] + list[i][index+2:]
					errCounter = true
				}
			}
		}
	}

	if errCounter {
		return dup(list)
	} else {
		return list
	}
}

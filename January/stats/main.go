package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// ************* INSTRUCTIONS ******************
// https://www.codewars.com/kata/55b3425df71c1201a800009c/train/go

// ************ EXAMPLE **********************
// "01|15|59, 1|47|6, 01|17|20, 1|32|34, 2|3|17"
//"Range: 00|47|18 Average: 01|35|15 Median: 01|32|34"

// range convert everything to mil, find the
//range and convert back to deserired format

func main() {
	// fmt.Println(past(1, 15, 59))
	// fmt.Println(past(1, 47, 6))
	// fmt.Println(past(1, 17, 20))
	// fmt.Println(past(1, 32, 34))
	// fmt.Println(past(2, 3, 17))
	solve("01|15|59, 1|47|6, 01|17|20, 1|32|34, 2|3|17")
}

func solve(s string) {

	strArr := strings.Split(s, ",")
	intArr := convertToMil(strArr)
	listRange := findRange(intArr)
	avarage := avarage(intArr)
	median := median(intArr)
	fmt.Printf("Range: %v Average: %v Median: %v", formatHMS(listRange), formatHMS(avarage), formatHMS(median))

}

func formatHMS(num int) string {
	// 2.77778e-7 hour
	fmt.Println("Before", num)
	var h string
	var s string
	var m string

	totalSeconds := num / 1000
	hour := math.Floor(float64(totalSeconds) / 60 / 60)
	if hour < 10 {
		h = fmt.Sprintf("0%v", hour)
	} else {
		h = fmt.Sprintf("%v", hour)
	}
	seconds := totalSeconds % 60
	if seconds < 10 {
		s = fmt.Sprintf("0%v", seconds)
	} else {
		s = fmt.Sprintf("%v", seconds)
	}

	minutes := math.Floor(float64(totalSeconds) / 60)
	if minutes > 60 {
		minutes = minutes - 60
	}

	if minutes < 10 {
		m = fmt.Sprintf("0%v", minutes)
	} else {
		m = fmt.Sprintf("%v", minutes)
	}

	out := fmt.Sprintf("%v|%v|%v", h, m, s)

	return out

}

func median(list []int) (out int) {
	sort.Ints(list)

	if len(list)%2 != 0 {
		out = list[len(list)-1] / 2
	}
	val1 := list[:len(list)/2]
	val2 := list[len(list)/2:]

	out = (val1[len(val1)-1] + val2[len(val2)-1]) / 2
	return out
}

func avarage(list []int) int {
	sum := 0
	for _, num := range list {
		sum += num
	}
	return sum / len(list)
}

func findRange(list []int) int {
	sort.Ints(list)
	return list[len(list)-1] - list[0]
}
func past(h, m, s int) int {
	return ((h * 3600) + (m * 60) + s) * 1000
}

func convertToMil(list []string) []int {
	out := make([]int, len(list))
	for i, item := range list {
		temp := strings.Split(item, "|")
		val1, err := strconv.Atoi(strings.TrimSpace(temp[0]))
		check(err)
		val2, err := strconv.Atoi(strings.TrimSpace(temp[1]))
		check(err)
		val3, err := strconv.Atoi(strings.TrimSpace(temp[2]))
		check(err)
		out[i] = past(val1, val2, val3)
	}
	return out
}

func check(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}

package main

// You get an array of numbers, return the sum of all of the positives ones.
// Example [1,-4,7,12] => 1 + 7 + 12 = 20
// Note: if there is nothing to sum, the sum is default to 0.

func main() {
	PostitiveSum([]int{1, -4, 7, 12})
}

//return the positive sum of all numbers that are positive
func PostitiveSum(numbers []int) int {

	count := 0

	for _, n := range numbers {

		if n > 0 {
			count += n
		}
	}

	return count

}

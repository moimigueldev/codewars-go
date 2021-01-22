package main

func main() {
	solve([]float32{1, 1, 1, 0.5, 1, 1})

}

func solve(arr []float32) float32 {

	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	for i, v := range arr[1:] {
		if v != arr[i] {
			return v
		}
	}
	return 0
}

package main

func IntToArray(number int, arr []int) []int {
	remainder := number
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = remainder % 10
		remainder /= 10
	}

	return arr
}
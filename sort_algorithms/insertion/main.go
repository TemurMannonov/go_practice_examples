package main

import "fmt"

// Time O(n^2) | Space O(1)
// Minimal O(n) when elements are already sorted
func main() {
	var array = []int{5, 3, 6, 7}
	insertionSort(&array)
	fmt.Println(array)
}

func insertionSort(array *[]int) {
	var i, j, key int
	for i = 1; i < len(*array); i++ {
		key = (*array)[i]
		j = i - 1
		for j >= 0 && key < (*array)[j] {
			(*array)[j+1] = (*array)[j]
			j--
		}
		(*array)[j+1] = key
	}
}

package main

import "fmt"

// Time O(n^2) | Space O(1)
// Minimal O(n) when elements are already sorted
func main() {
	var array = []int{5, 3, 67, 7}
	bubleSort(&array)
	fmt.Println(array)
}

func bubleSort(array *[]int) {
	var i, j int
	for i = 0; i < len(*array)-1; i++ {
		swapped := false
		for j = 0; j < len(*array)-i-1; j++ {
			if (*array)[j] > (*array)[j+1] {
				swap(&(*array)[j], &(*array)[j+1])
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}
}

func swap(number1, number2 *int) {
	temp := *number1
	*number1 = *number2
	*number2 = temp
}

package main

import "fmt"

// Time O(n^2) | Space O(1)
func main() {
	var array = []int{5, 3, 6, 7}
	selectionSort(&array)
	fmt.Println(array)
}

func selectionSort(array *[]int) {
	var i, j int
	for i = 0; i < len(*array); i++ {
		minIndex := i
		for j = i + 1; j < len(*array); j++ {
			if (*array)[j] < (*array)[minIndex] {
				minIndex = j
			}
		}
		swap(&(*array)[i], &(*array)[minIndex])
	}
}

func swap(number1, number2 *int) {
	temp := *number1
	*number1 = *number2
	*number2 = temp
}

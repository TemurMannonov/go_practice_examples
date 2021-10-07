package main

import "fmt"

func main() {
	var sumOdd, sumEven int

	for i := 1; i <= 100; i++ {
		if i % 2 != 0 {
			sumOdd += i
		} else {
			sumEven += i
		}
	}
	fmt.Printf("Sum of even: %d\n", sumEven)
	fmt.Printf("Sum of odd: %d\n", sumOdd)
	
}

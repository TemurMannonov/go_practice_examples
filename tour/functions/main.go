package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))


	//closures
	fmt.Println("Closures")
	pos, neg := adder(), adder()
	pos(5)
	pos(5)
	pos(5)
	fmt.Println(pos(5))
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	
}

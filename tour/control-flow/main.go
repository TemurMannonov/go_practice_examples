package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	var newValue float64
	for i := 1; i <= 10; i++ {
		newValue -= (z * z - x) / (2 * z)
		if (z - newValue) == 0 || math.Abs(z - newValue) <= 0.000000000000001 {
			break
		}
		z = newValue
		fmt.Println(z)
	}
	return z

}


func main() {

	//defer statement
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("defer started")



	//if statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)



	//Exercise
	fmt.Printf("\nSquare root\n")
	fmt.Println(Sqrt(81))


	//Exercise
	fmt.Printf("\nSwitch statement\n")
	fmt.Printf("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
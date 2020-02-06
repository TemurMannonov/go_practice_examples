package main

import (
	"fmt"
	"math"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	newValue := 1.0
	for i := 1; i <= 10; i++ {
		newValue -= (z*z -x) / (2 * z)
		if  (z - newValue) == 0 || math.Abs(z-newValue) <= 0.000000000000001 		 {
			break
		}
		z = newValue
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
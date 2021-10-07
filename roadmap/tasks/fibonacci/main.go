package main

import "fmt"

func fibonacci() func() int {
	f1, f2 := 1, 0
	return func() int {
		fib := f2
		f2, f1 = f1, fib + f1
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

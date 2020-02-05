package main

import "fmt"

func main() {
	var n, r, t int
	rev := 0

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	t = n
	for n > 0 {
		r = n % 10
		rev = rev * 10 + r
		n /= 10 
	}

	if t == rev {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Not palindrom")
	}
	
}

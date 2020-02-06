package main

import "fmt"

func Palindrom(s string) bool {
	l := len(s)
	for i := 0; i <= l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	p := Palindrom("abba")
	
	if p {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Not palindrom")
	}
	
	// var n, r, t int
	// rev := 0

	// fmt.Print("Enter a number: ")
	// fmt.Scan(&n)

	// t = n
	// for n > 0 {
	// 	r = n % 10
	// 	rev = rev * 10 + r
	// 	n /= 10 
	// }

	// if t == rev {
	// 	fmt.Println("Palindrom")
	// } else {
	// 	fmt.Println("Not palindrom")
	// }
	
}

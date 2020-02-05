package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 5
	y = sum - 5
	return 
}
func main() {
	fmt.Println(split(17))
}
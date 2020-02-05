package main

import "fmt"
import "strings"
// import "golang.org/x/tour/pic"

func main() {
	//array
	var a[2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a)


	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
	
	q := []int{2,3,4,5}
	fmt.Println(q)

	q = append(q, 100, 200, 300, 400, 500, 600, 700)
	fmt.Println(q)

	st := []struct{
		i int
		b bool
	}{
		{2, false},
		{5, true},
	}
	fmt.Println(st)

	printSlice(q)


	//dynamically-size array
	a1 := make([]int, 5)
	printSlice1("a1", a1)

	b1 := make([]int, 0, 5)
	printSlice1("b1", b1)

	//Tic Tac Toe
	board :=  [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}


	//range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}


	//Exercise
	// fmt.Println("\nExercise\n\n")
	// pic.Show(Pic)
}

// func Pic(dx, dy int) [][]uint8 {
// 	a := make([][]uint8, dy)
// 	for i := 0; i < dy; i++ {
// 		a[i] = make([]uint8, dx)
// 	}

// 	// Do something.
// 	for i := 0; i < dy; i++ {
// 		for j := 0; j < dx; j++ {
// 			switch {
// 			case j % 15 == 0:
// 				a[i][j] = 240
// 			case j % 3 == 0:
// 				a[i][j] = 120
// 			case j % 5 == 0:
// 				a[i][j] = 150
// 			default:
// 				a[i][j] = 100
// 			}
// 		}
// 	}	
// 	return a
// }

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
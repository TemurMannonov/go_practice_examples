package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 10 
	fmt.Println(v)

	p := &v
	p.X = 20
	fmt.Println(v)

	v1 := Vertex{X: 20, Y: 40}
	fmt.Println(v1)
}
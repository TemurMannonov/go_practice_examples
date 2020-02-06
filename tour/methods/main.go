package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	ScaleFunc(&v, 2)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)



	f := MyFloat(-4.56)
	fmt.Println(f.Abs())


	//using interface
	var a Abser
	k := MyFloat(-math.Sqrt2)
	t := Vertex{3, 4}

	a = k
	fmt.Println(a.Abs())
	a = &t
	fmt.Println(a.Abs())


	var i I = T{"hello"}
	i.M()
}


type Vertex struct {
	X, Y float64
}

//methods
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

//pointer receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


//function pointer
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


type MyFloat float64

//method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var n = map[string]Vertex {
	"Qarchi": Vertex{
		123.345, 2345.345,
	},
	"Kitob": Vertex{
		37.345, -123.34,
	},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.34, 34.12,
	}

	
	fmt.Println(m["Bell Labs"])
	fmt.Println(n)
}

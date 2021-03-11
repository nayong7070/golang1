package main

import . "fmt"

type Vertex struct {
	X, Y int
}

func (Vert *Vertex) mul() int {
	return Vert.X * Vert.Y
}

func main() {
	vert := Vertex{10, 20}
	Println(vert.mul())
	vert.X = 5
	vert.Y = 10
	Println(vert.mul())
}

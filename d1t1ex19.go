package main

import . "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.55852, -74.394858,
	}
	Println(m["Bell Labs"])
}

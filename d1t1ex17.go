package main

import . "fmt"

func main() {
	a := []int{1, 11, 111, 1111}
	b := [][]int{{1, 11, 111},
		{2, 22, 222},
		{3, 33, 333},
	}
	Println(a[0], a[1])
	Println(b[0], b[1], b[2][2])
}

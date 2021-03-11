package main

import . "fmt"

func main() {
	sum := 0
	addr := func(x int) int {
		sum += x
		return sum
	}
	aum := 1
	neg := func(x int) int {
		aum *= x
		return aum
	}
	for i := 1; i < 10; i++ {
		Println(addr(i), neg(i))
	}
}

package main

import (
	"fmt"
)

func main() {
	a := "Korea"[:]
	b := "Korea"[1:3]
	c := "Korea"[:2]
	d := "Korea"[2:]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

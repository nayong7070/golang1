package main

import (
	"fmt"
)

func main() {
	for i, ch := range "Korea 한국" {
		fmt.Printf("%d:%q", i, ch)
	}
}

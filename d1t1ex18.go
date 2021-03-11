package main

import . "fmt"

func main() {
	imap := map[int]string{
		0: "Notrh",
		1: "East",
		2: "West",
		3: "South",
	}
	Println(imap)
	Println(imap[1], imap[2])
}

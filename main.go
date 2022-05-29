package main

import (
	"fmt"
)

func main() {
	var stri = new(string)
	// *stri = "Hello Dude"
	if *stri == "" {
		fmt.Print("Empty")
	}
	fmt.Print(*stri)
}

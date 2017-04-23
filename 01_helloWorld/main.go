package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	for i := 0; i <= 200; i++ {
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i, i)
	}
}

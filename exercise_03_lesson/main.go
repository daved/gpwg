package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n = 42
	var g = -1

	for g != n {
		g = rand.Intn(100)
		switch {
		case g > n:
			fmt.Printf("the guess of %d is too high\n", g)
		case g < n:
			fmt.Printf("the guess of %d is too low\n", g)
		}
	}

	fmt.Printf("the guess of %d is just right\n", g)
}

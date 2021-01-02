package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var bank float64

	for bank < 20 {
		switch rand.Intn(3) {
		case 0:
			bank += 0.05
		case 1:
			bank += 0.10
		default:
			bank += 0.25
		}

		fmt.Printf("The bank contains $%2.2f\n", bank)
	}
}

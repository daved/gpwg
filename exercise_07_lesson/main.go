package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var bank int

	for bank < 2000 {
		switch rand.Intn(3) {
		case 0:
			bank += 5
		case 1:
			bank += 10
		default:
			bank += 25
		}

		fmt.Printf("The bank contains $%02d.%02d\n", bank/100, bank%100)
	}
}

package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	var year, month int

	for i := 0; i < 10; i++ {
		year = rand.Intn(4096)
		month = rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2:
			daysInMonth = 28
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				daysInMonth = 29
			}

		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1

		fmt.Println(era, year, month, day)
	}
}

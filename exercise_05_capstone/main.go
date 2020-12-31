package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const (
		minutelySeconds = 60
		hourlyMinutes   = 60
		hourlySeconds   = minutelySeconds * hourlyMinutes
		dailyHours      = 24
		dailySeconds    = hourlySeconds * dailyHours
		distance        = 62100000
	)

	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")

	for count := 0; count < 10; count++ {
		kms := rand.Intn(15) + 16
		days := distance / kms / dailySeconds
		price := 20.0 + kms

		var company string

		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		tripType := "One-way"
		if rand.Intn(2) == 1 {
			tripType = "Round-trip"
			price = price * 2
		}

		fmt.Printf("%-16s %4d %-10s $%4v\n", company, days, tripType, price)
	}
}

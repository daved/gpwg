package main

import "fmt"

func main() {
	const (
		dailyHrs = 24
		distance = 56000000.0
	)

	days := 28
	kph := distance / days / dailyHrs

	fmt.Printf("%d days to Malacandra at %v kph\n", days, kph)
}

package main

import "fmt"

func main() {
	const (
		distance      = 236e15
		lightSpeed    = 299792
		secondsToYear = 3154e4
		lightYears    = distance / lightSpeed / secondsToYear
	)

	fmt.Printf("The galaxy 'Canis Major Dwarf' is %f light years away.\n", lightYears)
}

package main

import "fmt"

func main() {
	s := "yes"

	switch s {
	case "true", "yes", "1":
		fmt.Println(true)
	case "false", "no", "0":
		fmt.Println(false)
	default:
		fmt.Printf("value not understood: %q\n", s)
	}
}

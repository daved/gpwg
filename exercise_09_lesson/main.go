package main

import "fmt"

func main() {
	caesar()
	international()
}

func caesar() {
	const (
		msg = "L fdph, L vdz, L frqtxhuhg."
	)

	for _, r := range msg {
		switch {
		case r >= 'a' && r <= 'z':
			r -= 3
			if r < 'a' {
				r += 26
			}
		case r >= 'A' && r <= 'Z':
			r -= 3
			if r < 'A' {
				r += 26
			}
		}
		fmt.Print(string(r))
	}
	fmt.Println()
}

func international() {
	const (
		msg = "Hola Estación Espacial Internacional"
	)

	for _, r := range msg {
		switch {
		case r >= 'a' && r <= 'z':
			r -= 13
			if r < 'a' {
				r += 26
			}
		case r >= 'A' && r <= 'Z':
			r -= 13
			if r < 'A' {
				r += 26
			}
		}
		fmt.Print(string(r))
	}
	fmt.Println()
}

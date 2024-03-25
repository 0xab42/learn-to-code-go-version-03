package main

import "fmt"

func main() {
	for x := 0; x < 5; x++ {
		fmt.Printf("#### x: %v ####\n", x)

		for y := 0; y < 5; y++ {
			fmt.Printf("x: %v; y: %v\n", x, y)
		}

		fmt.Printf("#### x: %v ####\n", x)
	}
}

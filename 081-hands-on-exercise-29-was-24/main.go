package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for x := 0; x < 100; x++ {
		fmt.Println(x)
	}

	for j := 0; j < 100; j++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("x: %v, y: %v\n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("both less than 4")
		case x > 6 && y > 6:
			fmt.Println("both gt 6")
		case x >= 4 && x <= 6:
			fmt.Println("x btwn 4-6")
		case y != 5:
			fmt.Println("y ne 5")
		default:
			fmt.Println("none...")
		}
	}
}

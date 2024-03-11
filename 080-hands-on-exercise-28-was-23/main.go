package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("x: %v, y: %v\n\n", x, y)

	// if x < 4 && y < 4 {
	// 	fmt.Println("both less than 4")
	// } else if x > 6 && y > 6 {
	// 	fmt.Println("both greater than 6")
	// } else if x >= 4 && x <= 6 {
	// 	fmt.Println("x between 4 and 6")
	// } else if y != 5 {
	// 	fmt.Println("y is not equal 5")
	// } else {
	// 	fmt.Println("none...")
	// }

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

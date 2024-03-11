package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for y := 1; y < 43; y++ {
		fmt.Printf("y: %v,\t", y)
		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Println("x = 0")
		case 1:
			fmt.Println("x = 1")
		case 2:
			fmt.Println("x = 2")
		case 3:
			fmt.Println("x = 3")
		case 4:
			fmt.Println("x = 4")
		default:
			fmt.Println("ERR: x = ", x)
		}
	}
}

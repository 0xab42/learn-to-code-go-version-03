package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)

	fmt.Printf("x: %v\t%T\n", x, x)

	if x >= 0 && x <= 100 {
		fmt.Println("x: between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("x: between 101 and 200")
	} else {
		fmt.Println("x: between 201 and 250")
	}
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var c int = 0
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			c++
			fmt.Printf("total: %v i=%v x=%v\n", c, i, x)
		}
	}
}

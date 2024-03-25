package main

import "fmt"

func main() {
	var x int = 100
	for {
		if x <= 0 {
			break
		}

		if x%2 == 0 {
			x--
			continue
		}

		fmt.Println(x)
		x--
	}
}

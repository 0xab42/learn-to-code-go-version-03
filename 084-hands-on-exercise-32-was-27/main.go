package main

import "fmt"

func main() {
	var i int = 0
	for {
		if i >= 100 {
			break
		}

		if i%2 == 0 {
			fmt.Println(i)
		}

		i++
	}
}

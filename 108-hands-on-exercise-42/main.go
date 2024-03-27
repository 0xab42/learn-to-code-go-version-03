package main

import "fmt"

func main() {
	var xa [5]int

	xa[0] = 1
	xa[1] = 2
	xa[2] = 3
	xa[3] = 4
	xa[4] = 5

	for i, v := range xa {
		fmt.Printf("index:%v, value: %v\n", i, v)
	}
}

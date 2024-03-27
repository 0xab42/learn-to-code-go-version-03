package main

import "fmt"

func main() {
	var xs []int

	for i, t := 0, 42; i < 10; i, t = i+1, t+1 {
		xs = append(xs, t)
	}

	for i, v := range xs {
		fmt.Printf("index: %v, type: %T, value: %v\n", i, v, v)
	}
}

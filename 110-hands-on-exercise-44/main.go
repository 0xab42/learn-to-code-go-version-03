package main

import "fmt"

func main() {
	var xs []int

	for i, t := 0, 42; i < 10; i, t = i+1, t+1 {
		xs = append(xs, t)
	}

	fmt.Println(xs[:5])
	fmt.Println(xs[5:])
	fmt.Println(xs[2:7])
	fmt.Println(xs[1:6])
}

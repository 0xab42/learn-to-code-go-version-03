package main

import "fmt"

func main() {
	m := map[string]int{
		"j": 42,
		"m": 32,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

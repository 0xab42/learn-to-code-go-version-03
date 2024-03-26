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

	m1 := m["j"]
	fmt.Println(m1)

	m2 := m["Q"]
	fmt.Println(m2)

	if m3, ok := m["Q"]; !ok {
		fmt.Println("Q not stored in the map 'm', and m3 is zero valued", m3)
	}

	if age, ok := m["j"]; ok {
		fmt.Println("age of 'j' is", age)
	}
}

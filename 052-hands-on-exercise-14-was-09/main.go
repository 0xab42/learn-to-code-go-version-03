package main

import "fmt"

var animal = "kot"

const name = "kokos"

func main() {
	k := 42
	fmt.Printf("value of k: %v, the type of k: %T\n", k, k)
	fmt.Printf("value of animal: %v, the type of animal: %T\n", animal, animal)
	fmt.Printf("value of name: %v, the type of name: %T\n", name, name)
}

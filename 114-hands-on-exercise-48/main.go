package main

import "fmt"

func main() {
	j1 := []string{"James", "Bond", "Shaken, not stirred"}
	j2 := []string{"Miss", "Moneypenny", "I'm 008."}

	jj := [][]string{j1, j2}

	for i, v := range jj {
		fmt.Printf("index: %v\nslice: %v\n\n", i, v)
	}

	for _, v := range jj {
		for _, iv := range v {
			fmt.Printf("%v\n", iv)
		}
	}
}

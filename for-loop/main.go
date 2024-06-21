package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}

	index := 10

	for index < 15 {
		fmt.Printf("%d ", index)
		index++
	}
}

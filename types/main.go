package main

import "fmt"

func main() {
	var i interface{} = 3.13

	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float32:
	case float64:
		fmt.Println("i is a float", i)
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}

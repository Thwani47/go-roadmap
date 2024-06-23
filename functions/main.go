package main

import "fmt"

var (
	area = func(l, b int) int {
		return l * b
	}
)

func rectangle(l int, b int) (perimeter int, area int) {
	perimeter = 2 * (l + b)
	area = l * b
	return
}

func sum(a, b int) int {
	return a + b
}

func partialSum(a int) func(int) int {
	return func(b int) int {
		return sum(a, b)
	}
}

func printColors(colors ...string) {
	fmt.Printf("%v\n", colors)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	func(l, b int) {
		fmt.Println(l * b)
	}(20, 10)
	perimeter, a := rectangle(20, 23)
	fmt.Printf("Perimeter: %d, Area: %d\n", perimeter, a)

	var l, b = 10, 20

	func() {
		fmt.Println(l * b)
	}()

	fmt.Println(area(20, 23))

	partial := partialSum(3)
	fmt.Println(partial(7))

	printColors()
	printColors("red")
	printColors("red", "green", "blue")

	fmt.Println(factorial(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

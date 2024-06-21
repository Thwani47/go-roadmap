package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var toBe bool = false
	var maxInt uint64 = 1<<64 - 1
	var z = cmplx.Sqrt(-5 + 2i)

	fmt.Printf("toBe is of type %T\n", toBe)
	fmt.Printf("maxInt is of type %T\n", maxInt)
	fmt.Printf("z is of type %T\n", z)

	// runes
	rune1 := 'B'
	rune2 := 's'
	rune3 := '\t'

	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %T\n", rune1, rune1, rune1)
	fmt.Printf("Rune 2: %c; Unicode: %U; Type: %T\n", rune2, rune2, rune2)
	fmt.Printf("Rune 3: %c; Unicode: %U; Type: %T\n", rune3, rune3, rune3)
}

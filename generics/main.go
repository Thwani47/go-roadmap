package main

import "fmt"

type Number interface {
	int | float64
}

// sumNumbers adds together the values of m
func sumNumbers[V Number](m map[string]V) V {
	var sum V

	for _, v := range m {
		sum += v
	}

	return sum
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	head, tail *Node[T]
}

func (l *List[T]) Push(v T) {
	if l.tail == nil {
		l.head = &Node[T]{Value: v}
		l.tail = l.head
	} else {
		l.tail.Next = &Node[T]{Value: v}
		l.tail = l.tail.Next
	}
}

func (l *List[T]) ListNodes() []T {
	var nodes []T

	for n := l.head; n != nil; n = n.Next {
		nodes = append(nodes, n.Value)
	}

	return nodes
}

func main() {
	ints := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	floats := map[string]float64{
		"one":   1.1,
		"two":   2.2,
		"three": 3.3,
	}

	fmt.Println(sumNumbers(ints))
	fmt.Println(sumNumbers(floats))

	var m = map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(MapKeys(m))

	var m2 = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(MapKeys(m2))

	l := List[int]{}
	l.Push(10)
	l.Push(20)
	l.Push(30)
	fmt.Println(l.ListNodes())

	l2 := List[string]{}
	l2.Push("one")
	l2.Push("two")
	l2.Push("three")
	fmt.Println(l2.ListNodes())
}

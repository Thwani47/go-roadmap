## Generics in Go
- Starting with Go 1.18, Go has added support for generics, also known as `type parameters`
- If we have these two functions
```go
// sumInts adds together the values of m
func sumInts(m map[string]int) int {
	var sum int

	for _, v := range m {
		sum += v
	}

	return sum
}

// sumFloats adds together the values of m
func sumFloats(m map[string]float64) float64 {
	var sum float64

	for _, v := range m {
		sum += v
	}

	return sum
}
```
we can create a generic function that works with both `int` and `float64` types
```go
// sum adds together the values of m
func sum[V int | float64](m map[string]V) V {
    var sum V

    for _, v := range m {
        sum += v
    }

    return sum
}
```
We can also define a custom type that can be either `int` or `float64`
```go
type Number interface {
    int | float64
}
```
And we can update our function as follows
```go
// sum adds together the values of m
func sum[V Number](m map[string]V) V {
    var sum V

    for _, v := range m {
        sum += v
    }

    return sum
}
```

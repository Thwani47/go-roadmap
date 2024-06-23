## Functions in Go
- We can name the return values in our functions
```go
func rectangle(l int, b int) (perimeter int, area int){
    perimeter = 2 * (l + b)
    area = l * b
    return
}

func main(){
    perimeter, area := rectangle(20, 23);
    fmt.Printf("Perimeter: %d, Area: %d\n", perimeter, area)
}
```
- We can write anonymous functions in Go. Anonymous functions are functions that are declared with no name
```go
var area = func(l, b int) int {
    return l * b
}

func main(){
    func(l, b int){
        fmt.Println(l * b)
    }(20, 10)

    fmt.Println(area(20, 10))
}
```
- We can also define closures in Go. Closures are anonymous functions that access the variables outside the function body
```go
func main(){
    var l, b = 10, 20
    func(){
        area := l * b
        fmt.Println(area)
    }
}
```
- We can also define high-order functions in Go. High-order functions are functions that either receive other functions as input or return other functions as output
- We can also define variadic functions in Go. Variadic functions are functions that can accept a variable number of arguments
```go
func printColors(colors ...string){
    // colors is a slice
    fmt.Println(colors)
}

func main(){
    printColors("red")
    printColors("red", "green", "blue")
}
```
- We can pass different variable types to a variadic function by defining the function as follows
```go
func doSomething(p ...interface{})
```
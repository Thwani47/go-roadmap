## Go Types, Type Assertions, and Switches
- A type assertion provides access to an interface value's underlying concrete value.
```go
t := i.(T)
```
This asserts that the interface value  `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.
For example:
```go
var i interface{} = 3.14

f := i.(float64)
fmt.Println(f) // 3.14
```
This asserts that `i` holds a `float64` value and assigns that value to `f`.
- If `i` does not hold a `float64`, the statement will trigger a panic.
- In order to test wheter `i` holds type `T`, an assertion can return two values, the value, and a boolean value that represents whether the assertion succeeded or not
```go
var i interface{} = 3.14

f, ok := i.(float64)

if !ok{
    fmt.Println("i is not a float64")
    return
}

fmt.Println(f) // 3.14
```
- A type switch is a construct that permits several type assertions in series
- It is like a regular switch statement, but the cases in a type switch specify types and not values. 
```go
var i interface{} = 3.14

switch v := i.(type){
case int:
    fmt.Println("i is an int")
case float64:
    fmt.Println("i is a float64")
case string:
    fmt.Println("i is a string")
default:
    fmt.Println("i is of another type")
}
```
The variable `v` will have the corresponding value of the type that matches the interface value `i`. i.e, in the above example, `v` will be of type `float64` and will have the value `3.14`.
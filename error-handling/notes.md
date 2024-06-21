## Error handling in Go
- In Go, errors are returned like any other type
- The `error` type is a built-in interface
```go
type error interface{
    Error() string
}
```
An error is anything that implements the `Error` method
- We can create a new error using the `errors` package
```go
import "errors"

func main(){
    err := errors.New("This is an error")
    fmt.Println(err)
}
```
- We can also use the `fmt.Errorf` function to create an error. This allows us to add dynamic data to the error message
```go
import "fmt"

func main(){
    err := fmt.Errorf("This is an error: %v", 42)
    fmt.Println(err)
}
```
- The zero-value of an error is `nil`
- We can use `errors.Is` to check if an error is of a certain type
- We can use `fmt.Errorf` and `%w` to wrap an error
```go
import "fmt"

func doSomething() error {
    // some logic 
    return fmt.Errorf("doSomething %w", errors.New("something went wrong"))
}

func doAnotherThing() error {
    // some logic 
    return fmt.Errorf("doAnotherThing %w", errors.New("an error occured"))
}

func main(){
    err := doSomething()
    fmt.Println(err)

    err = doAnotherThing()
    fmt.Println(err)
}
```
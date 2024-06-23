## Panic and Recover
### panic
- A panic is a mechanism that allows us to halt the normal execution of a program when an unexpected or unrecoverable situation occurs.
- A panic is similar to an exception in other languages
- When a panic occurs, the program starts to unwind the call stack, any deferred function calls are executed, and then crashes with a log message describing the cause of the panic
- The signature of the panic is 
```go
func panic(v interface{})
```
### recover
- The recover function is used to regain control of a panicking goroutine.
- We can use the recover function to catch a panic and resume normal execution.
- The recover function, when called inside a deferred function, returns the value that was passed to the panic function.
- The signature of the recover function is 
```go
func recover() interface{}
```
 
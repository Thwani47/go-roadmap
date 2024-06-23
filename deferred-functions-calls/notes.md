## Go Defferred function calls
- Go has a `defer` statement that schedules a function call to be run after the function completes.
- A defer statement is often used with paired operations like open and close, connect and disconnect, lock and unlock, to ensure that resources are released in all cases. 
- Without `defer`
```go
func doSomething() error{
    file.Open("somefile")

    if failureX {
        file.Close()
        return errX
    }

    if failureY {
        file.Close()
        return errY
    }

    file.Close()
    return nil
}
```
- With defer
```go
func doSomething() error{
    file.Open("somefile")
    defer file.Close()

    if failureX {
        return errX
    }

    if failureY {
        return errY
    }
    return nil
}
```
- The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
- Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
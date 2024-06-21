## Go For Loop

- Go has only one loop construct, the `for-loop`
- `for init; condition; post statement`
- We can omit the `init` and the `post statement`
```go
count = 0

for count < 10 {
    fmt.Println(count)
    count++
}
```
This is an equivalent of a `while loop` in other languages.
- If we omit the `init`, the `condition`, and the `post statement`, we get an infinite loop
```go
for {
    // do something
}
```
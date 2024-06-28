## Go Interfaces
- An interface type is defined as a set of method signatures
- A value of interface type hold any value that implements those methods
- A type implements an interface by implementing its methods
- There's no explicit `implements` keyword
- Interface values can be thought of as a tuple of a value and a concrete type: `(value, type)`
- If the concrete value inside the interface itself is `nil`, the method will be called with a `nil` receiver
- A nil interface value holds neither value nor concrete type. Calling a method on a nil interface throws a run-time error because there is no type inside the interface tuple to indicate which concrete method to call
```go
type I interface {
    Method()
}

func main(){
    var i I
    fmt.Prinf("(%v, %y)\n", i, i) // <nil>, <nil>
    i.Method() // run-time error: invalid memory address or nil pointer dereference
}
```
- The interface type that has 0 methods is known as the `empty interface`, `interface{}`
- An empty interface may hold values of any type. Every type implements at least zero methods
- Empty interfaces are used by code that handles values of unknown type
## Go Modules
- A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root
- The `go.mod` defines the module's module path, which is also the import path used for the root directory, and its dependency requirements
- The `go.mod` only appears at the root of the module
- Run this command to list a module and all its dependencies
```bash
go list -m all
```
- Run this command to list all available versions of a module
```bash
go list -m -versions <module-name>
```
- Run this command to view the docs of a module
```bash
go doc <module-name>
```
- `go mod tidy` helps remove unused dependencies
- `go build` and `go test` can only tell you about missing dependencies, but cannot remove unused dependecies.
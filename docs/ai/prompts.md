# Prompts

## 2026-05-31 `package` declaration in Go files

**Prompt**
Explain the `package` declaration at the top of Go files.

**Response Summary**:
Every single Go file must start with a `package` declaration. It defines the **namespace** for that file.

- **Shared Scope**: All files in the same directory with the same package name can see and use each other's functions, variables and types **without importing them**.
- **Encapsulation**: If a variable or function starts with a lowercase letter (eg. `func calculateTotal()`) it is only visible to other files withing that *same* package. If it starts with a capital letter (eg. `func InitWindow()`) it can be exported and can be used by other packages.

`package main` tells the compiler that this package should compile to an executable instead of a share library. It must contain a `func main()`, which serves as an entry point of the application.

### Multiple Command Line Utilities

If we want to create multiple executables withing the same project (each with it's own `package main` and `func main()` ), we cannot place them in the same directory. The Go standard convention to to put each utility in it's own directory inside `cmd/`. For eg:

```
my-project/
├── go.mod
└── cmd/
    ├── server/
    │   └── main.go  (package main - builds the server)
    └── cli/
        └── main.go  (package main - builds the CLI tool)
```
This layout is followed in the current repo which has `cmd/goboxd/main.go`.
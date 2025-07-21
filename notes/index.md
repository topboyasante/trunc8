# Introduction: Trunc8

## Packages and Modules

### Packages

In go, a package is a folder that groups related `.go` files together. All the `.go` files in that folder belong to the same package.

**Key rules:**

- Every .go file must start with a package declaration
- All files in the same directory must have the same package name
- The special package main is where your program starts (needs a main() function)

**How importing works:**

- You import packages by their path: import "myproject/internal/handlers"
- You can only use **exported functions/types** from other packages. Exported functions/types are ones that **start with a capital letter.**

Example

```
url-shortener/
├── cmd/server/main.go        (package main)
├── internal/handlers/        (package handlers)
├── internal/models/          (package models)
└── internal/database/        (package database)
```

### Modules

In go, a modules is a collection of [packages](#packages).

When you run `go mod init <project_name>`, a module is initialized, and a `go.mod` file is created.

**What the go.mod file tracks:**

- Your module's name/path (what you just created)
- What Go version you're using
- External packages your project depends on (like npm dependencies)
- The specific versions of those dependencies

**Why modules matter:**

- Dependency management: When you import external packages (like a PostgreSQL driver), Go automatically adds them to go.mod
- Versioning: Go can handle different versions of the same package
- Reproducible builds: Anyone can clone your project and get the exact same dependency versions.

**Simple example:**

When you eventually add the PostgreSQL driver with something like:

`go import "github.com/lib/pq"`

Go will automatically update your go.mod file to include that dependency and create a `go.sum` file (like package-lock.json) that locks the exact version.

"Locking the exact version" means ensuring everyone gets the identical code, down to every single line. It stores a cryptographic hash (like a fingerprint) of the exact code you downloaded.


Next: [Project Structure](/notes/project-structure.md).
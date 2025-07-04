# Go ptr utility

[![Go Reference](https://pkg.go.dev/badge/github.com/denisdubochevalier/ptr.svg)](https://pkg.go.dev/github.com/denisdubochevalier/ptr)

A tiny, generic Go utility package to create pointers from literal values

## The problem

In Go, you cannot take the address of a literal or a constant value directly. This means code like the following is invalid and will not compile:

```go
// This code will NOT compile
config := APIConfig{
  Timeout: &30,         // Error: cannot take the address of 30
  Retries: &5,          // Error: cannot take the address of 5
  Enabled: &true,       // Error: cannot take the address of true
  DefaultUser: &"admin" // Error: cannot take the address of "admin"
}
```

The standard workaround is to declare an intermediate variable for each value, which can be verbose and clutter your code.

```go
// This is the standard, but verbose, way
tmpTimeout := 30
tmpRetries := 5
tmpEnabled := true
tmpUser := "admin"

config := APIConfig{
  Timeout: &tmpTimeout,
  Retries: &tmpRetries,
  Enabled: &tmpEnabled,
  DefaultUser: &tmpUser,
}
```

## The solution

This `ptr` package provides a single, generic helper function, `ptr.New`, to solve this problem elegantly. It allows you to create a pointer from any value in a single, clean expression.

```go
import "github.com/denisdubochevalier/ptr"

// This is the clean, concise way using this package
config := APIConfig{
  Timeout: ptr.New(30),
  Retries: ptr.New(5),
  Enabled: ptr.New(true),
  DefaultUser: ptr.New("admin"),
}
```

### Why use pointers for fields?

Using a pointer for a field (e.g., `*int`, `*bool`) is a common pattern in Go for distinguishing between a **zero value** (like `0`, `false`, or an empty string `""`) and a value that was **not provided** (`nil`). This is especially useful in:

- **Configuration**: Determining if a setting was intentionally set to its zero value or omitted entirely.
- **API Payloads (JSON/XML)**: Correctly handling optional fields that might be omitted from a request.

## Installation

```bash
go get github.com/denisdubochevalier/ptr
```

## Usage

Simply import the package and use `ptr.New()` with any literal value.

```go
package main

import (
  "fmt"
  "github.com/denisdubochevalier/ptr"
)

type ServerConfig struct {
  Port *int
  Host *string
  Verbose *bool
}

func main() {
  // Example: Initializing a struct with pointer fields
  config := ServerConfig{
    Port:    ptr.New(8080),
    Host:    ptr.New("localhost"),
    Verbose: ptr.New(true),
  }

  if config.Port != nil {
    fmt.Printf("Server port: %d\n", *config.Port)
  }

  if config.Host != nil {
    fmt.Printf("Server host: %s\n", *config.Host)
  }
}
```

## Testing

To run the unit tests for this package, clone the repository and run:

```bash
go test -v -race ./...
```

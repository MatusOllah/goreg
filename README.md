# 📋 goreg

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/goreg.svg)](https://pkg.go.dev/github.com/MatusOllah/goreg) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/goreg)](https://goreportcard.com/report/github.com/MatusOllah/goreg) [![Go](https://github.com/MatusOllah/goreg/actions/workflows/go.yml/badge.svg)](https://github.com/MatusOllah/goreg/actions/workflows/go.yml) [![GitHub license](https://img.shields.io/github/license/MatusOllah/goreg)](https://github.com/MatusOllah/goreg/blob/main/LICENSE) [![Made in Slovakia](https://raw.githubusercontent.com/pedromxavier/flag-badges/refs/heads/main/badges/SK.svg)](https://www.youtube.com/watch?v=UqXJ0ktrmh0)

**goreg** is a lightweight, thread-safe, and generic registry package for Go.

## Features

* Thread safety using mutexes
* Go generics support
* JSON and Gob serialization
* Ordered registries for order-sensitive values
* Easy-to-use methods for registration and retrieval

## Use Cases

* Application-wide configuration management
* Plugin systems
* Game modloaders and object registration

## Installation

Run:

```sh
go get -u github.com/MatusOllah/goreg
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/MatusOllah/goreg"
)

type Thing string

func main() {
    // Create a new empty registry
    reg := goreg.NewStandardRegistry[Thing]()

    // Register a thing
    reg.Register("door", Thing("Door"))

    // Retrieve the thing
    thing, ok := reg.Get("door")
    if ok {
        fmt.Println("Thing found:", thing)
    } else {
        fmt.Println("Thing not found")
    }

    // Unregister the thing
    reg.Unregister("door")
}
```

### `StandardRegistry[T]` vs. `OrderedRegistry[T]`

There are 2 types of registries in goreg: The `StandardRegistry[T]` and the `OrderedRegistry[T]`.

`StandardRegistry[T]` is a standard registry that uses a map under the hood. Output order is non-deterministic due to Go's map implementation. You can use this for most things that don't require specific order (e.g. users, clients).

`OrderedRegistry[T]` is however an **ordered registry**. It uses a slice of key-value pairs under the hood. You can use this for things that **require specific order** (e.g. game levels, chapters).

Example code:

```go
package main

import (
    "fmt"

    "github.com/MatusOllah/goreg"
)

type Level struct {
    Name string
    // ...
}

func main() {
    stdReg := goreg.NewStandardRegistry[Level]()
    stdReg.Register("level1", Level{Name: "Level 1"})
    stdReg.Register("level2", Level{Name: "Level 2"})
    stdReg.Register("level3", Level{Name: "Level 3"})
    stdReg.Register("level4", Level{Name: "Level 4"})

    fmt.Println(goreg.Collect(stdReg)) // Output order is unpredictable due to Go maps, should print out in random order

    ordReg := goreg.NewOrderedRegistry[Level]()
    ordReg.Register("level1", Level{Name: "Level 1"})
    ordReg.Register("level2", Level{Name: "Level 2"})
    ordReg.Register("level3", Level{Name: "Level 3"})
    ordReg.Register("level4", Level{Name: "Level 4"})

    fmt.Println(goreg.Collect(ordReg)) // Should print out correctly in order
}
```

## License

Licensed under the **MIT License** (see [LICENSE](https://github.com/MatusOllah/goreg/blob/main/LICENSE))

# go-deepcopy

[![Go Reference](https://pkg.go.dev/badge/github.com/iqbalmind/go-deepcopy.svg)](https://pkg.go.dev/github.com/iqbalmind/go-deepcopy)
[![Go Report Card](https://goreportcard.com/badge/github.com/iqbalmind/go-deepcopy)](https://goreportcard.com/report/github.com/iqbalmind/go-deepcopy)

A simple reflection-based **deep copy library for Go**, supporting:
- Structs
- Pointers
- Slices
- Maps
- Arrays
- Interfaces
- Channels

## 🚀 Installation

```bash
go get github.com/iqbalmind/go-deepcopy
```

## 📖 Usage
```go
package main

import (
    "fmt"
    "github.com/iqbalmind/go-deepcopy"
)

type Address struct {
    Street string
    City   string
}

type Person struct {
    Name    string
    Age     int
    Address *Address
}

func main() {
    original := Person{
        Name: "Dewi",
        Age:  28,
        Address: &Address{
            Street: "Jl. Braga No. 99",
            City:   "Bandung",
        },
    }

    cloned, _ := deepcopy.DeepCopy(original)
    clonedPerson := cloned.(Person)

    // Modify the cloned struct
    clonedPerson.Name = "Budi"
    clonedPerson.Address.Street = "Jl. Asia Afrika No. 45"

    fmt.Println("Original:", original.Name, original.Address.Street)
    fmt.Println("Cloned:", clonedPerson.Name, clonedPerson.Address.Street)
}

```

## Output:
```bash
Original: Dewi Jl. Braga No. 99
Cloned:   Budi Jl. Asia Afrika No. 45
```

## 🧪 Running Tests
``go
go test -v
``

## 📜 License
MIT © 2025 [iqbalmind](https://github.com/iqbalmind)
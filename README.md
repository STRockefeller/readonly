# Go Package `readonly`

## Overview

The `readonly` package provides a simple and efficient way to handle immutable and final values in Go. It offers two main types: `readOnly` and `final`, each designed to manage values that are either immutable or assigned once.

## Installation

To use the `readonly` package in your Go project, simply import it as follows:

```go
import "github.com/STRockefeller/readonly"
```

## Usage

### The `readOnly` Struct

The `readOnly` struct is used to create immutable values. Once a value is set, it cannot be changed.

#### Creating a `readOnly` Instance

```go
value := readonly.NewReadOnly(42)
```

#### Accessing the Value

```go
v := value.Get()
```

### The `final` Struct

The `final` struct allows for creating a value that can be set only once. It's an alternative to `readOnly` for scenarios where the value might not be known at the time of instantiation.

#### Creating a `final` Instance

```go
finalValue := readonly.NewFinal[int]()
```

#### Setting the Value

To set the value for the first time:

```go
err := finalValue.Set(42)
if err != nil {
    // handle error
}
```

#### Accessing the Value

To safely access the value:

```go
v, err := finalValue.Get()
if err != nil {
    // handle error
}
```

To access the value with a panic on error:

```go
v := finalValue.MustGet()
```

## Features

- Immutable value handling with `readOnly`.
- Single-assignment value handling with `final`.
- Type-safe implementation using Go generics.

## Error Handling

Both `final.Get()` and `final.Set()` methods return errors to handle situations where the value is not set or being reassigned, respectively. Use `MustGet()` and `MustSet()` for panicking versions.

## License

[MIT License](./LICENSE)

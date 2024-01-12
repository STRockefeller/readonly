# Go Package `readonly`

## Overview

The `readonly` package provides a simple and efficient way to handle immutable and final values in Go. It offers two main types: `ReadOnly` and `Final`, each designed to manage values that are either immutable or assigned once.

## Installation

To use the `readonly` package in your Go project, simply import it as follows:

```go
import "github.com/STRockefeller/readonly"
```

## Usage

### The `ReadOnly` Struct

The `ReadOnly` struct is used to create immutable values. Once a value is set, it cannot be changed.

#### Creating a `ReadOnly` Instance

```go
value := readonly.NewReadOnly(42)
```

#### Accessing the Value

```go
v := value.Get()
```

### The `Final` Struct

The `Final` struct allows for creating a value that can be set only once. It's an alternative to `ReadOnly` for scenarios where the value might not be known at the time of instantiation.

#### Creating a `Final` Instance

```go
var finalValue readonly.Final[int]
```

#### Setting the Value

To set the value for the first time:

```go
finalValue.Set(42)
```

#### Accessing the Value

To safely access the value:

```go
v := finalValue.Get()
```

## Features

- Immutable value handling with `ReadOnly`.
- Single-assignment value handling with `Final`.
- Type-safe implementation using Go generics.

## License

[MIT License](./LICENSE)

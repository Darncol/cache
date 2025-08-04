# Go Simple Cache test project

[![Go Reference](https://pkg.go.dev/badge/github.com/Darncol/cache.svg)](https://pkg.go.dev/github.com/Darncol/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/Darncol/cache)](https://goreportcard.com/report/github.com/Darncol/cache)

A simple in-memory cache implementation in Go with thread-safe operations.

## Features

- Simple key-value storage
- Thread-safe operations
- Support for any value type using `interface{}`
- Basic CRUD operations
- Zero dependencies

## Installation

```bash
go get github.com/Darncol/cache
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/Darncol/cache"
)

func main() {
	// Create a new cache instance
	c := cache.New()

	// Set values
	tc.Set("name", "John Doe")
	tc.Set("age", 30)

	// Get values
	name := tc.Get("name")
	fmt.Println("Name:", name) // Output: Name: John Doe

	// Delete a key
	tc.Delete("age")

	// Try to get deleted key
	age := tc.Get("age")
	fmt.Println("Age:", age) // Output: Age: <nil>
}
```

## API

### `New() *Cache`
Creates a new cache instance.

### `(c *Cache) Set(key string, value interface{})`
Sets a value in the cache with the specified key.

### `(c *Cache) Get(key string) interface{}`
Retrieves a value from the cache by key. Returns `nil` if key doesn't exist.

### `(c *Cache) Delete(key string)`
Removes a key-value pair from the cache.

## License

MIT

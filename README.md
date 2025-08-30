# Go Simple Cache test project

[![Go Reference](https://pkg.go.dev/badge/github.com/Darncol/cache.svg)](https://pkg.go.dev/github.com/Darncol/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/Darncol/cache)](https://goreportcard.com/report/github.com/Darncol/cache)

A simple in-memory cache implementation in Go with thread-safe operations and item expiration.

## Features

- Simple key-value storage
- Thread-safe operations
- Support for any value type using `interface{}`
- Item expiration
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
	"time"
	"github.com/Darncol/cache"
)

func main() {
	// Create a new cache instance
	c := cache.New()

	// Set values with expiration
	c.Set("name", "John Doe", 5 * time.Second)
	c.Set("age", 30, 10 * time.Second)

	// Get values
	name, exists := c.Get("name")
	if exists {
		fmt.Println("Name:", name) // Output: Name: John Doe
	}

	// Wait for "name" to expire
	time.Sleep(6 * time.Second)

	name, exists = c.Get("name")
	if !exists {
		fmt.Println("Name after expiration: not found") // Output: Name after expiration: not found
	}

	// Delete a key before it expires
	c.Delete("age")

	// Try to get deleted key
	age, exists := c.Get("age")
	if !exists {
		fmt.Println("Age: not found") // Output: Age: not found
	}
}
```

## API

### `New() *Cache`
Creates a new cache instance.

### `(c *Cache) Set(key string, value interface{}, duration time.Duration)`
Sets a value in the cache with the specified key and expiration duration. If the key already exists, its value and timer will be updated.

### `(c *Cache) Get(key string) (interface{}, bool)`
Retrieves a value from the cache by key. Returns the value and a boolean indicating whether the key was found.

### `(c *Cache) Delete(key string)`
Removes a key-value pair from the cache and stops its expiration timer.

## License

MIT
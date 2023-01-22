# cache-ninja
In-memory key-value storage

```
go get -u github.com/ninja-way/cache-ninja
```

```go
// New return cache pointer
func New() *cache.Cache

// Cache is a struct for storage any value by unique key
type Cache struct {}
    // Set add to storage new value by key 
    Set(key string, value interface{})
    // Get return stored value by key if it exists
    Get(key string) (interface{}, error)
    // Delete remove value from storage by key if it exists
    Delete(key string) error
```

### Example

```go
package main

import (
	"fmt"
	"github.com/ninja-way/cache-ninja/pkg/cache"
)

func main() {
	cache := cache.New()

	cache.Set("id", 1)

	fmt.Println(cache.Delete("user"))

	cache.Set("id", 2)
	cache.Set("user", struct {
		name string
		age  int8
	}{
		name: "Tom",
		age:  19,
	})

	_, err := cache.Get("test")
	if err != nil {
		fmt.Println(err)
	}

	value, _ := cache.Get("id")
	fmt.Println(value)
}
```
#### Output
```
delete: unknown key
get: unknown key
2  
```

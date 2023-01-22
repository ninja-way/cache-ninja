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
    // Set add new value by key to storage for a specific time ttl
    Set(key string, value interface{}, ttl time.Duration)
    // Get return stored value by key if it exists, else returns an error
    Get(key string) (interface{}, error)
    // Delete remove value from storage by key if it exists, else returns an error
    Delete(key string) error
```

### Example

```go
package main

import (
	"fmt"
	"github.com/ninja-way/cache-ninja/pkg/cache"
	"time"
)

func main() {
	cache := cache.New()

	// Try get non-exist id
	_, err := cache.Get("test")
	if err != nil {
		fmt.Println(err)
	}
	// Result: get: unknown key

	// Try delete non-exist id
	fmt.Println(cache.Delete("user"))
	// Result: delete: unknown key

	// Set and then get value
	cache.Set("user_name", "Trevor", time.Second*3)
	value, _ := cache.Get("user_name")
	fmt.Println(value)
	// Result: Trevor

	// Add items in goroutines
	for i := 1; i <= 100; i++ {
		duration := time.Second
		if i%2 == 0 {
			duration = time.Minute
		}

		go cache.Set(fmt.Sprintf("%did", i), i, duration)
	}
	time.Sleep(time.Second * 2)

	// Get paired with big ttl and unpaired with small ttl
	fmt.Println(cache.Get("21id"))
	fmt.Println(cache.Get("22id"))
	// Result: <nil> get: unknown key
	//	    22 <nil>
}
```
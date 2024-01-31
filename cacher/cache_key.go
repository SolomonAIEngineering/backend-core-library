package cacher // import "github.com/SolomonAIEngineering/backend-core-library/cacher"

import "fmt"

type CacheKey string

// NewCacheKey creates a new CacheKey with the given key.
func NewCacheKey(key string) CacheKey {
	return CacheKey(key)
}

// String returns the string representation of the cache key.
func (c CacheKey) String() string {
	return string(c)
}

// Enrich enriches the cache key with the given value.
func (c CacheKey) Enrich(value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%s:%s", c.String(), value)
}

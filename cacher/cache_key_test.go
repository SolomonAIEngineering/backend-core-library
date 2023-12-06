package cacher

import (
	"testing"
)

func TestCacheKey_Enrich(t *testing.T) {
	c := NewCacheKey("test-cache-key") // Declare and initialize the variable c

	tests := []struct {
		name  string
		c     CacheKey
		value string
		want  string
	}{
		{
			name:  "Empty value",
			c:     c,
			value: "",
			want:  "",
		},
		{
			name:  "Non-empty value",
			c:     c,
			value: "example",
			want:  "test-cache-key:example",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.Enrich(tt.value)
			if got != tt.want {
				t.Errorf("CacheKey.Enrich() = %v, want %v", got, tt.want)
			}
		})
	}
}

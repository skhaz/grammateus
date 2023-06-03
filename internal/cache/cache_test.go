package cache_test

import (
	"testing"
	"time"

	"skhaz.dev/streamopinion/internal/cache"
)

func TestSetAndGet(t *testing.T) {
	var (
		cache         = cache.NewCache()
		key           = "key"
		expectedValue = struct{ data string }{data: "test"}
	)

	cache.Set(key, expectedValue, time.Second)
	value, _ := cache.Get(key)

	if value != expectedValue {
		t.Errorf("Expected value be %v, but got %v", expectedValue, value)
	}
}

func TestSetAndExpire(t *testing.T) {
	var (
		cache         = cache.NewCache()
		key           = "key"
		expectedValue = struct{ data string }{data: "test"}
	)

	cache.Set(key, expectedValue, time.Second)
	time.Sleep(2 * time.Second)
	_, ok := cache.Get(key)

	if ok {
		t.Errorf("Expected item be expired")
	}
}

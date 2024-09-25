package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestReapLoopConcurrent(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Add("https://example.com", []byte("testdata"))
			_, ok := cache.Get("https://example.com")
			if !ok {
				t.Errorf("expected to find key")
			}
		}()
	}

	wg.Wait()

	time.Sleep(waitTime)

	_, ok := cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
	}
}

func TestAddEdgeCases(t *testing.T) {
	cache := NewCache(5 * time.Second)

	// Test empty key
	key := ""
	val := []byte("testdata")
	cache.Add(key, val)
	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	// Test nil value
	key = "https://example.com"
	val = nil
	cache.Add(key, val)
	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected to not find key")
		return
	}

	// Test duplicate keys
	key = "https://example.com"
	val = []byte("testdata")
	cache.Add(key, val)
	cache.Add(key, []byte("newtestdata"))
	val, ok = cache.Get(key)
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	if string(val) != "newtestdata" {
		t.Errorf("expected to find new value")
		return
	}
}

func TestConcurrentAccess(t *testing.T) {
	cache := NewCache(10 * time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
			cache.Add(key, []byte("testdata"))
			_, ok := cache.Get(key)
			if !ok {
				t.Errorf("expected to find key")
			}
		}(fmt.Sprintf("key-%d", i))
	}

	wg.Wait()
}

package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://example.com",
			value: []byte("testdata"),
		},
		{
			key:   "https://example.com/path",
			value: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.value)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.value) {
				t.Errorf("value in cache doesn't match expected")
				return
			}
		})
	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	var testKey string = "https://example.com"

	cache := NewCache(baseTime)
	cache.Add(testKey, []byte("testdata"))

	_, ok := cache.Get(testKey)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(testKey)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

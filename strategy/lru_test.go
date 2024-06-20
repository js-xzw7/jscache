package strategy

import "testing"

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	cache := NewLruCache(int64(0), nil)

	cache.Add("key1", String("123456"))

	if v, ok := cache.Get("key1"); !ok || string(v.(String)) != "123456" {
		t.Fatalf("cache hit key1=1234 failed")
	}

	if _, ok := cache.Get("key2"); !ok {
		t.Fatalf("cache not key2")
	}
}

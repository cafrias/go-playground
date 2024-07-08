package lrucache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := LRUCache[string, int]{
		lookup:        make(map[string]*node[int]),
		reverseLookup: make(map[*node[int]]string),
		capacity:      3,
	}

	if _, ok := lru.Get("foo"); ok {
		t.Fatal(`'foo' should not be in the cache`)
	}

	lru.Upsert("foo", 69)
	if val, _ := lru.Get("foo"); val != 69 {
		t.Fatal(`'foo' should be in the cache`)
	}

	lru.Upsert("bar", 420)
	if val, _ := lru.Get("bar"); val != 420 {
		t.Fatal(`'bar' should be in the cache`)
	}

	lru.Upsert("baz", 1337)
	if val, _ := lru.Get("baz"); val != 1337 {
		t.Fatal(`'baz' should be in the cache`)
	}

	lru.Upsert("ball", 69420)
	if val, _ := lru.Get("ball"); val != 69420 {
		t.Fatal(`'ball' should be in the cache`)
	}

	if _, ok := lru.Get("foo"); ok {
		t.Fatal(`'foo' must have been evicted`)
	}

	if _, ok := lru.Get("bar"); !ok {
		t.Fatal(`'bar' should be in the cache`)
	}

	lru.Upsert("foo", 69)

	if _, ok := lru.Get("bar"); !ok {
		t.Fatal(`'bar' should be in the cache`)
	}

	if _, ok := lru.Get("foo"); !ok {
		t.Fatal(`'foo' should be in the cache`)
	}

	if _, ok := lru.Get("baz"); ok {
		t.Fatal(`'baz' must have been evicted`)
	}
}

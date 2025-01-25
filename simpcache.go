package simpcache

import (
	"sync"
	"time"
)

const NoExp time.Duration = -1
const DefExp time.Duration = 0

type Item[Val any] struct {
	value   Val
	expires int
}

type cache[k ~string, V any] struct {
	mu         sync.RWMutex
	items      map[k]*Item[V]
	done       chan struct{}
	expTime    time.Duration
	cleanupInt time.Duration
}

type Cache[K ~string, V any] struct {
	*cache[K, V]
}

func newCache[K ~string, V any](expTime, cleanupInt time.Duration, Item map[K]*Item[V]) *cache[K, V] {
	c := &cache[K, V]{
		mu:         sync.RWMutex{},
		items:      Item,
		done:       make(chan struct{}),
		expTime:    expTime,
		cleanupInt: cleanupInt,
	}
	return c
}

func New[K ~string, V any](expTime, cleanupInt time.Duration) *Cache[K, V] {

	Items := make(map[K]*Item[V])
	c := newCache(expTime, cleanupInt, Items)

	return &Cache[K, V]{c}
}

func (c *Cache[K, V]) add(key K, value V, d time.Duration) error {
	var exp int64

	if d == DefExp {
		d = c.expTime
	}
	if d > 0 {
		exp = time.Now().Add(d).UnixNano()
	}
	if d < 0 {
		exp = int64(NoExp)
	}
	c.mu.Lock()
	c.items[key] = &Item[V]{
		value:   value,
		expires: int(exp),
	}
	return nil
}

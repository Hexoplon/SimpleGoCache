package simplegocache

import (
	"time"

	"github.com/pkg/errors"
)

// MemCache - memory cache object, implements Cache interface
type MemCache struct {
	TTL            int64
	MaxElements    int
	MaxElementSize int
	Entries        map[string]Entry
}

// Entry - A cache entry
type Entry struct {
	Content []byte
	Created int64
}

// NewMemCache - Initiates a new memory based cache
func NewMemCache(ttl int64, maxElements int, maxElementSize int) *MemCache {
	return &MemCache{
		TTL:            ttl,
		MaxElements:    maxElements,
		MaxElementSize: maxElementSize,
		Entries:        make(map[string]Entry, maxElements),
	}
}

// AddToCache - Add []byte to cache with key
func (c *MemCache) AddToCache(content []byte, key string) error {
	if c.InCache(key) {
		return errors.New(ErrorAlreadyExists)
	}
	if c.toLarge(&content) {
		return errors.New(ErrorTooLarge)
	}

	if c.toLong() {
		return errors.New(ErrorTooManyEntries)
	}

	c.Entries[key] = Entry{
		Content: content,
		Created: time.Now().Unix(),
	}

	return nil
}

// ReadFromCache - Read the value of a given key from cache
func (c *MemCache) ReadFromCache(key string) ([]byte, error) {
	if c.InCache(key) {
		return c.Entries[key].Content, nil
	}
	return nil, errors.New(ErrorNoEntry)
}

// InCache - Check if a key is registered in cache
func (c *MemCache) InCache(key string) bool {
	if _, ok := c.Entries[key]; ok {
		if c.expired(key) {
			c.DeleteFromCache(key)
			return false
		}
		return true
	}
	return false
}

// DeleteFromCache - Delete an object from cache
func (c *MemCache) DeleteFromCache(key string) {
	delete(c.Entries, key)
}

// UpdateCache - Update the contents of a key already in cache
func (c *MemCache) UpdateCache(content []byte, key string) error {
	if c.toLarge(&content) {
		return errors.New(ErrorTooLarge)
	}

	if c.toLong() {
		return errors.New(ErrorTooManyEntries)
	}

	c.Entries[key] = Entry{
		Content: content,
		Created: time.Now().Unix(),
	}

	return nil
}

// Prune - prunes items that have expired
func (c *MemCache) Prune() {
	for k := range c.Entries {
		if c.expired(k) {
			c.DeleteFromCache(k)
		}
	}
}

// Close - Remove cached files
func (c *MemCache) Close() {
	for k := range c.Entries {
		c.DeleteFromCache(k)
	}
}

// expired - checks if a cache entry has expired
func (c *MemCache) expired(key string) bool {
	if c.Entries[key].Created+c.TTL < time.Now().Unix() {
		return true
	}
	return false
}

func (c *MemCache) toLarge(content *[]byte) bool {
	return len(*content) > c.MaxElementSize
}

func (c *MemCache) toLong() bool {
	return len(c.Entries) >= c.MaxElements
}

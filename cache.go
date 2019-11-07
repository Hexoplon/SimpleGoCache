package simplegocache

/*

Implements file based caching

A new cache is created with NewCache(folder, ttl, prefix). Remember to close with c.Close()!

To add content to cache use AddToCache(content, key) with content as a string and a string key.

To read content from cache use ReadFromCache(key). This returns content as a string.

To check whether a key exists in cache use InCache(key). Returns a bool

To delete a key use DeleteFromCache(key).

To Update a cache entry use UpdateCache(content, key)
a
*/

// Cache - cache interface to support different backend implementations. A cache should also have a constructor
type Cache interface {
	AddToCache()
	ReadFromCache()
	InCache()
	DeleteFromCache()
	UpdateCache()
	Prune()
	Close()
	expired()
}

// CacheStore - Stores and manages caches
type CacheStore struct {
}

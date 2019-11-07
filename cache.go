package simplegocache

import (
	"context"
)

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

type CacheStore struct {
	Caches map[string]*MemCache
}

func (c CacheStore) NewCache(ctx context.Context, req *NewCacheMsg, rsp *CacheMsg) error {
	if _, ok := c.Caches[req.Key]; !ok {
		c.Caches[req.Key] = NewMemCache(req.Ttl, int(req.MaxElements), int(req.MaxElementSize))
		rsp.Msg = MsgOk
	} else {
		rsp.Msg = MsgNotOk
	}

	return nil
}

func (c CacheStore) Add(ctx context.Context, req *EntryMsg, rsp *CacheMsg) error {
	return nil
}

func (c CacheStore) Delete(ctx context.Context, req *EntryMsg, rsp *Empty) error {
	panic("implement me")
}

func (c CacheStore) Read(ctx context.Context, req *EntryMsg, rsp *EntryMsg) error {
	panic("implement me")
}

func (c CacheStore) InCache(ctx context.Context, req *EntryMsg, rsp *EntryMsg) error {
	panic("implement me")
}

func (c CacheStore) Update(ctx context.Context, req *EntryMsg, rsp *CacheMsg) error {
	panic("implement me")
}

func (c CacheStore) Prune(ctx context.Context, req *Empty, rsp *Empty) error {
	panic("implement me")
}

func (c CacheStore) Close(ctx context.Context, req *Empty, rsp *Empty) error {
	panic("implement me")
}

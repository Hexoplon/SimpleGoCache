package simplegocache

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
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

// Cache - Contains references to cached objects
type Cache struct {
	CacheFolder string
	TTL         int64
	Prefix      string
	Files       map[string]File
}

// File - Contains info on a particular cached file
type File struct {
	Path    string
	TTL     int64
	Created int64
}

// NewCache - Initiates a new cache
func NewCache(folder string, ttl int64, prefix string) (*Cache, error) {
	var tmpFolder string
	tmpFolder, err := ioutil.TempDir(folder, prefix)
	if err != nil {
		log.Println("Unable to create temp folder for cache: " + err.Error())
	}
	return &Cache{CacheFolder: tmpFolder, TTL: ttl, Prefix: prefix, Files: make(map[string]File)}, nil
}

// Close - Remove cached files
func (c *Cache) Close() {
	err := os.RemoveAll(c.CacheFolder)
	if err != nil {
		log.Println("Unable to remove tmp cache folder: " + err.Error())
	}
}

// AddToCache - Add []byte to cache with key
func (c *Cache) AddToCache(content []byte, key string) error {
	// check that file does not exist
	if c.InCache(key) {
		return errors.New("Key" + key + " already exists in map!")
	}

	tmpfile, err := ioutil.TempFile(c.CacheFolder, c.Prefix)
	if err != nil {
		fmt.Println("Encountered error while writing file: " + err.Error())
		return err
	}

	c.Files[key] = File{
		Path:    tmpfile.Name(),
		TTL:     c.TTL,
		Created: time.Now().Unix(),
	}

	_, err = tmpfile.Write(content)
	if err != nil {
		return err
	}

	return nil
}

// ReadFromCache - Read the value of a given key from cache
func (c *Cache) ReadFromCache(key string) ([]byte, error) {
	if c.InCache(key) {
		content, err := ioutil.ReadFile(c.Files[key].Path)
		if err != nil {
			return []byte{}, err
		}
		return content, nil
	}

	return []byte{}, errors.New("Cannot find key: " + key + " in cache!")
}

// InCache - Check if a key is registered in cache
func (c *Cache) InCache(key string) bool {
	if _, ok := c.Files[key]; ok {
		if (time.Now().Unix() - c.Files[key].Created) > c.Files[key].TTL {
			delete(c.Files, key)
		}

		return true
	}

	return false
}

// DeleteFromCache - Delete an object from cache
func (c *Cache) DeleteFromCache(key string) error {
	if _, ok := c.Files[key]; ok {
		err := os.Remove(c.Files[key].Path)
		if err != nil {
			return err
		}
		delete(c.Files, key)
		return nil
	}

	return errors.New("key doesn't exist in cache")
}

// UpdateCache - Update the contents of a key already in cache
func (c *Cache) UpdateCache(content []byte, key string) error {
	err := c.DeleteFromCache(key)
	if err != nil {
		return err
	}
	err = c.AddToCache(content, key)
	if err != nil {
		return err
	}

	return nil
}

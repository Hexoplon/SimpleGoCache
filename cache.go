package SimpleGoCache

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

*/

type Cache struct {
	CacheFolder string
	TTL         int64
	Prefix      string
	Files       map[string]File
}

type File struct {
	Path    string
	TTL     int64
	Created int64
}

func NewCache(folder string, TTL int64, prefix string) (*Cache, error) {
	var tmpFolder string
	tmpFolder, err := ioutil.TempDir(folder, prefix)
	if err != nil {
		log.Println("Unable to create temp folder for cache: " + err.Error())
	}
	return &Cache{CacheFolder: tmpFolder, TTL: TTL, Prefix: prefix, Files: make(map[string]File)}, nil
}

func (c *Cache) Close() {
	err := os.RemoveAll(c.CacheFolder)
	if err != nil {
		log.Println("Unable to remove tmp cache folder: " + err.Error())
	}
}

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

func (c *Cache) ReadFromCache(key string) ([]byte, error) {
	if c.InCache(key) {
		content, err := ioutil.ReadFile(c.Files[key].Path)
		if err != nil {
			return []byte{}, err
		}
		return content, nil
	} else {
		return []byte{}, errors.New("Cannot find key: " + key + " in cache!")
	}
}

func (c *Cache) InCache(key string) bool {
	if _, ok := c.Files[key]; ok {
		if (time.Now().Unix() - c.Files[key].Created) > c.Files[key].TTL {
			delete(c.Files, key)
		} else {
			return true
		}
	}

	return false
}

func (c *Cache) DeleteFromCache(key string) error {
	if _, ok := c.Files[key]; ok {
		err := os.Remove(c.Files[key].Path)
		if err != nil {
			return err
		}
		delete(c.Files, key)
		return nil
	} else {
		return errors.New("key doesn't exist in cache")
	}
}

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

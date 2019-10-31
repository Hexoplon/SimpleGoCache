package SimpleGoCache

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"git.gvk.idi.ntnu.no/zackeusb/assignment1"

	"gotest.tools/assert"
)

func Test_AddToCache(t *testing.T) {
	//create tmp string to cache
	content := "This is a cache add test"
	testKey := "test"

	//Create new cache
	cache, err := assignment1.NewCache("", 60, "test")
	//Check for errors when creating cache
	assert.NilError(t, err)
	defer cache.Close()

	// Add content to cache
	err = cache.AddToCache([]byte(content), testKey)
	assert.NilError(t, err)

	//Test content of file
	tmpfnpath := cache.Files[testKey].Path
	readContent, err := ioutil.ReadFile(tmpfnpath)

	assert.Equal(t, string(readContent), content)
}

func Test_ReadFromCache(t *testing.T) {
	//setup cache
	//create tmp string to cache
	content := "This is a cache read test"
	testKey := "test"

	//Create new cache
	cache, err := assignment1.NewCache("", 60, "test")
	//Check for errors when creating cache
	assert.NilError(t, err)
	defer cache.Close()

	//Setup cache file
	tmpfnpath := filepath.Join(cache.CacheFolder, "testfile.cache")
	ioutil.WriteFile(tmpfnpath, []byte(content), 0666)

	// Manually add file to cache
	cache.Files[testKey] = assignment1.File{
		Path:    tmpfnpath,
		TTL:     60,
		Created: time.Now().Unix(),
	}

	//Attempt to read from cache
	readContent, err := cache.ReadFromCache(testKey)

	assert.Equal(t, string(readContent), content)

}

func Test_DeleteFromCache(t *testing.T) {
	testFiles := []assignment1.File{{
		Path:    "/tmp/test",
		TTL:     10,
		Created: time.Now().Unix(),
	},
		{
			Path:    "/tmp/test2",
			TTL:     20,
			Created: time.Now().Unix(),
		},
	}

	testKeys := []string{"a", "b"}

	cache, err := assignment1.NewCache("", 20, "test")
	assert.NilError(t, err)
	defer cache.Close()

	for i, v := range testFiles {
		cache.Files[testKeys[i]] = v
		ioutil.WriteFile(cache.Files[testKeys[i]].Path, []byte("This is a test file"), 0666)
	}

	//Delete first entry
	err = cache.DeleteFromCache(testKeys[0])
	assert.NilError(t, err)

	//Check that first entry is gone
	_, ok := cache.Files[testKeys[0]]
	assert.Assert(t, !ok)

	//Verify second entry still exists
	_, ok = cache.Files[testKeys[1]]
	assert.Assert(t, ok)

	//Cleanup
	for _, v := range cache.Files {
		os.Remove(v.Path)
	}

}

func Test_UpdateCache(t *testing.T) {
	//setup cache
	//create tmp string to cache
	content := "This is a cache read test"
	testKey := "test"

	//Create new cache
	cache, err := assignment1.NewCache("", 60, "test")
	//Check for errors when creating cache
	assert.NilError(t, err)
	defer cache.Close()

	//Setup cache file
	tmpfnpath := filepath.Join(cache.CacheFolder, "testfile.cache")
	ioutil.WriteFile(tmpfnpath, []byte(content), 0666)

	// Manually add file to cache
	cache.Files[testKey] = assignment1.File{
		Path:    tmpfnpath,
		TTL:     60,
		Created: time.Now().Unix(),
	}
	updatedContent := "This file has updated content"
	err = cache.UpdateCache([]byte(updatedContent), testKey)
	assert.NilError(t, err)

	readContent, err := ioutil.ReadFile(cache.Files[testKey].Path)

	assert.Equal(t, string(readContent), updatedContent)

}

func Test_InCache(t *testing.T) {
	testFiles := []assignment1.File{{
		Path:    "/tmp/test",
		TTL:     10,
		Created: time.Now().Unix(),
	},
		{
			Path:    "/tmp/test2",
			TTL:     0,
			Created: time.Now().Unix(),
		},
	}

	testKeys := []string{"a", "b"}

	cache, err := assignment1.NewCache("", 20, "test")
	assert.NilError(t, err)
	defer cache.Close()

	for i, v := range testFiles {
		cache.Files[testKeys[i]] = v
	}
	//Sleep to allow for 2nd entry to expire
	time.Sleep(2000000000)

	//Should exist
	assert.Assert(t, cache.InCache(testKeys[0]))
	//Should expire, be removed and not exist
	assert.Assert(t, !cache.InCache(testKeys[1]))
	//Should not exist
	assert.Assert(t, !cache.InCache("IdontExist"))
}

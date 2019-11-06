package simplegocache

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestMemCache_AddToCache(t *testing.T) {
	var testContent = []byte("This is some great test content")
	var testContentLen = len(testContent)
	var key = "test"

	memCache := NewMemCache(60, 3, testContentLen*2)

	err := memCache.AddToCache(testContent, key)
	assert.NilError(t, err)

	// Check for single entry
	_, ok := memCache.Entries[key]
	assert.Equal(t, ok, true)

	// Verify second entry is rejected
	err = memCache.AddToCache(testContent, key)
	assert.Error(t, err, ErrorAlreadyExists)

	// Check for too many items error
	err = memCache.AddToCache(testContent, key+"1")
	assert.NilError(t, err)

	err = memCache.AddToCache(testContent, key+"2")
	assert.NilError(t, err)

	err = memCache.AddToCache(testContent, key+"3")
	assert.Error(t, err, ErrorTooManyEntries)

	// Clean before next test
	delete(memCache.Entries, key)

	// Verify content size error
	err = memCache.AddToCache(append(append(testContent, testContent...), byte(0)), key)
	assert.Error(t, err, ErrorTooLarge)
}

func TestMemCache_Close(t *testing.T) {
	var testContent = []byte("This is some great test content")
	var testContentLen = len(testContent)
	var key = "test"

	memCache := NewMemCache(60, 3, testContentLen*2)

	// Setup test data
	memCache.Entries[key] = Entry{
		Content: testContent,
		Created: time.Now().Unix(),
	}
	memCache.Entries[key+"1"] = Entry{
		Content: testContent,
		Created: time.Now().Unix(),
	}

	memCache.Close()
	// Verify that both items are gone
	assert.DeepEqual(t, memCache.Entries[key], Entry{})
	assert.DeepEqual(t, memCache.Entries[key+"1"], Entry{})
}

func TestMemCache_DeleteFromCache(t *testing.T) {
	var testContent = []byte("This is some great test content")
	var testContentLen = len(testContent)
	var key = "test"

	memCache := NewMemCache(60, 3, testContentLen*2)

	// Setup test data
	memCache.Entries[key] = Entry{
		Content: testContent,
		Created: 10,
	}
	memCache.Entries[key+"1"] = Entry{
		Content: testContent,
		Created: 10,
	}

	memCache.DeleteFromCache(key)
	// Verify that one item is gone
	assert.DeepEqual(t, memCache.Entries[key], Entry{})
	assert.DeepEqual(t, memCache.Entries[key+"1"], Entry{Content: testContent, Created: 10})
}

func TestMemCache_InCache(t *testing.T) {
	var testContent = []byte("This is some great test content")
	var testContentLen = len(testContent)
	var key = "test"

	memCache := NewMemCache(60, 3, testContentLen*2)

	err := memCache.AddToCache(testContent, key)
	assert.NilError(t, err)

	// Verify invalid key does not exist
	assert.Equal(t, memCache.InCache("invalid"), false)

	// Verify valid key exists
	assert.Equal(t, memCache.InCache(key), true)
}

func TestMemCache_Prune(t *testing.T) {
	var testContent = []byte("This is some great test content")
	var testContentLen = len(testContent)
	var key = "test"
	var now = time.Now().Unix()

	memCache := NewMemCache(5, 3, testContentLen*2)

	// Setup test data
	memCache.Entries[key] = Entry{
		Content: testContent,
		Created: now + 10,
	}
	memCache.Entries[key+"1"] = Entry{
		Content: testContent,
		Created: now - 10,
	}

	memCache.Prune()
	// Verify that both items are gone
	assert.DeepEqual(t, memCache.Entries[key], Entry{Content: testContent, Created: now + 10})
	assert.DeepEqual(t, memCache.Entries[key+"1"], Entry{})
}

func TestMemCache_ReadFromCache(t *testing.T) {

}

func TestMemCache_UpdateCache(t *testing.T) {

}

func TestNewMemCache(t *testing.T) {

}

package simplegocache

import "testing"

var testContent = []byte("This is some great test content")
var testContentLen = len(testContent)

func TestMemCache_AddToCache(t *testing.T) {
	memCache := NewMemCache(60, 3, testContentLen*2)

}

func TestMemCache_Close(t *testing.T) {

}

func TestMemCache_DeleteFromCache(t *testing.T) {

}

func TestMemCache_InCache(t *testing.T) {

}

func TestMemCache_Prune(t *testing.T) {

}

func TestMemCache_ReadFromCache(t *testing.T) {

}

func TestMemCache_UpdateCache(t *testing.T) {

}

func TestNewMemCache(t *testing.T) {

}

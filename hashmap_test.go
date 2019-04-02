package hashmap

import (
	"testing"
)

func TestHashMapGet(t *testing.T) {
	var hashMap *HashMap
	hashMap = new(HashMap)
	hashMap.Init(3)

	_, bool := hashMap.Get("notexist")

	if bool != false {
		t.Fatalf("Expected %v but got %v", false, bool)
	}
}

func TestHashMapPutABool(t *testing.T) {
	var hashMap *HashMap
	hashMap = new(HashMap)
	hashMap.Init(3)

	hashMap.Put("hello", true)
	val, bool := hashMap.Get("hello")

	if bool != true {
		t.Fatalf("Expected %v but got %v", true, bool)
	}

	if val != true {
		t.Fatalf("Expected %v but got %v", true, val)
	}
}

func TestHashMapPutAString(t *testing.T) {
	var hashMap *HashMap
	hashMap = new(HashMap)
	hashMap.Init(3)

	hashMap.Put("hello", "world")
	val, bool := hashMap.Get("hello")

	if bool != true {
		t.Fatalf("Expected %v but got %v", true, bool)
	}

	if val != "world" {
		t.Fatalf("Expected %v but got %v", "world", val)
	}
}

func TestHashMapDel(t *testing.T) {
	var hashMap *HashMap
	hashMap = new(HashMap)
	hashMap.Init(3)

	hashMap.Put("hello", "world")
	currLen := hashMap.Len
	bool := hashMap.Del("hello")

	if bool != true {
		t.Fatalf("Expected %v but got %v", true, bool)
	}

	expectedLen := currLen - 1
	if expectedLen != hashMap.Len {
		t.Fatalf("Expected %v but got %v", expectedLen, hashMap.Len)
	}
}

func TestHashMapRehashing(t *testing.T) {
	var hashMap *HashMap
	hashMap = new(HashMap)
	hashMap.Init(3)

	currentCap := hashMap.Cap

	hashMap.Put("hello", true)

	if hashMap.Cap != 3 {
		t.Fatalf("Capacity should be unchanged, expected %v, got %v", 3, hashMap.Cap)
	}

	hashMap.Put("world", 3.14)
	hashMap.Put("foo", "bar")
	hashMap.Put("bar", "baz")

	expectedCap := currentCap * 2

	if hashMap.Cap != expectedCap {
		t.Fatalf("Capacity should be doubled, expected %v, got %v", expectedCap, hashMap.Cap)
	}
}
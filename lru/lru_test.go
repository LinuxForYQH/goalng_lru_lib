package main

import (
	"testing"
)

func TestCache_Common(t *testing.T) {
	cache := Cache{
		maxcap:     3,
		currentcap: 0,
		elementMap: make(map[interface{}]*ListElement),
	}
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Get(1)
	if cache.first.element.key != 1 {
		t.Error("first is error")
	}
	if cache.last.element.key != 2 {
		t.Error("last is error")
	}
}

func TestCache_FullDelet(t *testing.T) {
	cache := Cache{
		maxcap:     3,
		currentcap: 0,
		elementMap: make(map[interface{}]*ListElement),
	}
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	if _, i := cache.Get(1); i != 1 {
		t.Error("element is lose")
	}
	cache.Put(4, 4)
	if _, i := cache.Get(2); i == 1 {
		t.Error("element is not delete")
	}
	if _, i := cache.Get(4); i != 1 {
		t.Error("element is lose")
	}
}

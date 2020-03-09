package main

import "fmt"

func main() {
	cache := Cache{
		maxcap:     3,
		currentcap: 0,
		elementMap: make(map[interface{}]*ListElement),
	}
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	fmt.Println(cache.Get(1))
	cache.Put(4, 4)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(4))
}

type Cache struct {
	first      *ListElement
	last       *ListElement
	elementMap map[interface{}]*ListElement
	maxcap     int
	currentcap int
}

type Element struct {
	key interface{}
	val interface{}
}

type ListElement struct {
	pre     *ListElement
	element *Element
	next    *ListElement
}

func (c *Cache) Get(key interface{}) (interface{}, int) {
	if le, ok := c.elementMap[key]; ok {
		if le == c.first {
		} else if le == c.last {
			if le.pre != nil {
				le.pre.next = nil
				c.last = le.pre
			}
			le.pre = nil
			c.first.pre = le
			le.next = c.first
			c.first = le
		} else {
			pre := le.pre
			next := le.next
			pre.next = next
			next.pre = pre
			le.pre = nil
			le.next = nil
			c.first.pre = le
			le.next = c.first
			c.first = le
		}
		return le.element.val, 1
	}
	return nil, 0
}

func (c *Cache) Put(key, value interface{}) {
	//看一下是否存在,存在那么放第一
	if _, i := c.Get(key); i == 1 {
		c.first.element.key = value
		return
	}
	element := &Element{key: key, val: value}
	//移除队尾的一个
	if c.maxcap == c.currentcap {
		delete(c.elementMap, c.last.element.key)
		if c.last.pre != nil {
			c.last.pre.next = nil
		} else {
			c.last = nil
			c.first = nil
		}
		c.currentcap--
	}
	listElement := &ListElement{
		element: element,
	}
	c.elementMap[key] = listElement
	if c.first != nil {
		c.first.pre = listElement
		listElement.next = c.first
		c.first = listElement
	} else {
		c.first = listElement
	}
	if c.last == nil {
		c.last = listElement
	}
	c.currentcap++
}

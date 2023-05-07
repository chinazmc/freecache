package freecache

import (
	"container/list"
)

type firstLru struct {
	cap  int
	list *list.List
}

func newFirstLRU(size int) *firstLru {
	return &firstLru{
		cap:  size,
		list: list.New(),
	}
}

func (lru *firstLru) add(key []byte) (eitem []byte, evicted bool) {
	if lru.list.Len() < lru.cap {
		lru.list.PushFront(key)
		return nil, false
	}
	//如果容量满了
	evictItem := lru.list.Back()
	item := evictItem.Value.([]byte)
	return item, true
}

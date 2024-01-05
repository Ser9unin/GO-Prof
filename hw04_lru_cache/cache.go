package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	goroutineLock sync.Mutex
	capacity      int
	queue         List
	items         map[Key]*ListItem
}

type cacheItem struct {
	key Key
	val interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.goroutineLock.Lock()
	defer l.goroutineLock.Unlock()

	_, keyInCache := l.items[key]

	if keyInCache {
		l.queue.Remove(l.items[key])
	} else if l.queue.Len() >= l.capacity {
		itemToRemove := l.queue.Back().Value.(cacheItem)

		l.queue.Remove(l.items[itemToRemove.key])
		delete(l.items, itemToRemove.key)
	}

	l.items[key] = l.queue.PushFront(cacheItem{key: key, val: value})

	return keyInCache
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.goroutineLock.Lock()
	defer l.goroutineLock.Unlock()

	itemInCache, keyInCache := l.items[key]

	if !keyInCache {
		return nil, false
	}

	itemVal := itemInCache.Value.(cacheItem).val
	l.queue.MoveToFront(itemInCache)

	return itemVal, keyInCache
}

func (l *lruCache) Clear() {
	l.goroutineLock.Lock()
	defer l.goroutineLock.Unlock()

	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

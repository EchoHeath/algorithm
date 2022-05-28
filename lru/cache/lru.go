package cache

import "algorithm/lru/doubleLinkedList"

type LruCache struct {
	Capacity int                                    //总容量
	Size     int                                    //当前使用的容量
	LruMap   map[interface{}]*doubleLinkedList.Node //缓存
	LruList  *doubleLinkedList.DoubleLinkedList     //双向链表
}

func NewLruCacheService(capacity int) *LruCache {
	return &LruCache{
		Capacity: capacity,
		Size:     0,
		LruMap:   make(map[interface{}]*doubleLinkedList.Node, capacity),
		LruList:  doubleLinkedList.NewDoubleLinkedListService(capacity),
	}
}

//Get 每次获取将其放到链表头部
func (l *LruCache) Get(key interface{}) interface{} {
	if node, ok := l.LruMap[key]; ok {
		l.LruList.Remove(node)
		l.LruList.AppendFront(node)
		return node.Value
	}else {
		return -1
	}
}

//Put put keys.
func (l *LruCache) Put(key, value interface{}) {
	if l.Capacity == 0 {
		return
	}

	if node, ok := l.LruMap[key]; ok {
		l.LruList.Remove(node)
		node.Value = value
		l.LruList.AppendFront(node)
	}else {
		if l.Size == l.Capacity {
			tail := l.LruList.RemoveTail()
			delete(l.LruMap, tail.Key)
			l.Size--
		}

		node = &doubleLinkedList.Node{
			Key:   key,
			Value: value,
		}
		l.LruList.AppendFront(node)
		l.LruMap[key] = node
		l.Size++
	}
}

func (l *LruCache) Delete(key interface{}) {
	if l.Size == 0 {
		return
	}
	if node, ok := l.LruMap[key]; ok {
		l.LruList.Remove(node)
		delete(l.LruMap, key)
		l.Size--
	}
}

func (l *LruCache) Print() {
	l.LruList.Print()
}

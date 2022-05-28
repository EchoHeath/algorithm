package cache

import "algorithm/fifo/doubleLinkedList"

type FifoCache struct {
	Capacity int                                    //总容量
	Size     int                                    //当前使用的容量
	FifoMap  map[interface{}]*doubleLinkedList.Node //缓存
	FifoList *doubleLinkedList.DoubleLinkedList     //双向链表
}

func NewFifoCacheService(capacity int) *FifoCache {
	return &FifoCache{
		Capacity: capacity,
		FifoMap:  map[interface{}]*doubleLinkedList.Node{},
		FifoList: doubleLinkedList.NewDoubleLinkedListService(capacity),
	}
}

//Get get a cache.
func (f *FifoCache) Get(key interface{}) interface{} {
	if v, ok := f.FifoMap[key]; !ok {
		return -1
	} else {
		return v.Value
	}
}

//Put put a cache.
func (f *FifoCache) Put(key, value interface{}) {
	if f.Capacity == 0 {
		return
	}

	if node, ok := f.FifoMap[key]; ok {
		f.FifoList.Remove(node)
		node.Value = value
		f.FifoList.Append(node)
	} else {
		if f.Size == f.Capacity {
			head := f.FifoList.Pop()
			delete(f.FifoMap, head.Key)
			f.Size--
		}
		node = &doubleLinkedList.Node{
			Key:   key,
			Value: value,
		}
		f.FifoList.Append(node)
		f.FifoMap[key] = node
		f.Size++
	}
}

//Delete delete a list.
func (f *FifoCache) Delete(key interface{}) {
	if f.Size == 0 {
		return
	}
	if node, ok := f.FifoMap[key]; ok {
		f.FifoList.Remove(node)
		delete(f.FifoMap, key)
		f.Size--
	}
}

//Print print list.
func (f *FifoCache) Print() {
	f.FifoList.Print()
}

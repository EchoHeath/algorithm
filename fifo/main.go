package main

import (
	"algorithm/fifo/cache"
	"fmt"
)

func main() {
	lruCache := cache.NewFifoCacheService(2)
	lruCache.Put(1, 1)
	lruCache.Print()
	fmt.Printf("get a value: %v\n", lruCache.Get(1))
	lruCache.Put(2, 2)
	lruCache.Print()
	lruCache.Put(3, 3)
	fmt.Printf("get a value: %v\n", lruCache.Get(3))
	lruCache.Print()
	lruCache.Delete(2)
	lruCache.Print()
	lruCache.Delete(3)
	lruCache.Print()
	fmt.Printf("get a value: %v\n", lruCache.Get(3))
}

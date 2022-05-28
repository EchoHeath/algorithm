package main

import (
	"algorithm/lru/cache"
	"fmt"
)

func main() {
	lruCache := cache.NewLruCacheService(4)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)
	lruCache.Print()
	fmt.Printf("get a value: %v\n", lruCache.Get(1))
	lruCache.Print()
	lruCache.Put(5, 5)
	lruCache.Print()
	lruCache.Put(3, 3)
	fmt.Printf("get a value: %v\n", lruCache.Get(3))
	lruCache.Print()
	lruCache.Delete(1)
	lruCache.Print()

}

package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu    sync.Mutex
	items map[string]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		items: make(map[string]string),
	}
}

func (sm *SafeMap) Get(key string) (string, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, ok := sm.items[key]
	return value, ok
}

func (sm *SafeMap) Set(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.items[key] = value
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.items, key)
}

func main() {

	safeMap := NewSafeMap()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)

			safeMap.Set(key, value)

			retrievedValue, found := safeMap.Get(key)
			if found {
				fmt.Printf("Goroutine %d: Qiymat olingan: %s\n", i, retrievedValue)
			} else {
				fmt.Printf("Goroutine %d: Qiymat topilmadi\n", i)
			}

			safeMap.Delete(key)
			fmt.Printf("Goroutine %d: Qiymat o'chirildi\n", i)
		}(i)
	}

	wg.Wait()
}

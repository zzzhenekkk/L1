package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func (sm *SafeMap) Write(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Read(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.m[key]
	return val, ok
}

func main() {
	sm := SafeMap{m: make(map[string]int)}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			sm.Write(key, n)
			fmt.Printf("Written: %s = %d\n", key, n)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value, _ := sm.Read(key)
		fmt.Printf("Read: %s = %d\n", key, value)
	}
}

// ////// также можно реализовать используя sync.Map
func main2() {
	var sm sync.Map

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			sm.Store(key, n)
			fmt.Printf("Written: %s = %d\n", key, n)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		if value, ok := sm.Load(key); ok {
			fmt.Printf("Read: %s = %d\n", key, value)
		}
	}
}

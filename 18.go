package main

import (
	"fmt"
	"sync"
)

// counter - структура-счетчик
type Counter struct {
	mu    sync.Mutex
	value int
}

// increment - метод для инкрементирования счетчика
func (c *Counter) Increment() {
	c.mu.Lock()         // блокируем мьютекс
	defer c.mu.Unlock() // разблокируем мьютекс при выходе из функции
	c.value++           // увеличиваем значение счетчика
}

// value - метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()         // блокируем мьютекс
	defer c.mu.Unlock() // разблокируем мьютекс при выходе из функции
	return c.value      // возвращаем текущее значение счетчика
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// количество горутин
	numGoroutines := 100

	// запускаем горутины
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()     // сообщаем о завершении горутины
			counter.Increment() // инкрементируем счетчик
		}()
	}

	// ожидаем завершения всех горутин
	wg.Wait()

	// выводим итоговое значение счетчика
	fmt.Printf("final counter value: %d\n", counter.Value())
}

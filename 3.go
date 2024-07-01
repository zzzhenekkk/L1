package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений

func main() {
	var wg sync.WaitGroup

	arr := [...]int{2, 4, 6, 8, 10}
	var mu sync.Mutex

	var sum int = 0

	wg.Add(len(arr))
	for _, value := range arr {
		go func() {
			defer wg.Done()
			mu.Lock()
			sum += value * value
			mu.Unlock()
		}()

	}

	wg.Wait()
	println(sum)
}

func square(n int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- n * n
}

func main2() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, len(numbers))

	// Запускаем горутину для каждого числа в массиве
	for _, n := range numbers {
		wg.Add(1)
		go square(n, &wg, ch)
	}

	// Закрываем канал, когда все горутины завершены
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Суммируем результаты
	sum := 0
	for result := range ch {
		sum += result
	}

	fmt.Printf("The sum of squares is: %d\n", sum)
}

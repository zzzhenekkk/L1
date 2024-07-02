package main

import (
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

// 1 - х
// 2 - 2х

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	first := make(chan int)
	second := make(chan int)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()
		for _, value := range arr {
			first <- value
		}
		close(first)
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for value := range first {
			second <- value * value
		}
		close(second)
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for value := range second {
			println(value)
		}
	}()

	wg.Wait()

}

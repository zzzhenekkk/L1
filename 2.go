package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(len(arr))
	for i, value := range arr {
		go func() {
			defer wg.Done()
			arr[i] = value * value
		}()
	}
	wg.Wait()
	fmt.Println(arr)
}

package main

import (
	"fmt"
)

// binarySearch выполняет бинарный поиск элемента target в отсортированном массиве arr
// и возвращает индекс элемента или -1, если элемент не найден.
func binarySearch(arr []int, target int) (int, bool) {
	middle := len(arr) / 2

	if middle == 0 {

	}
	for i := 0; i < len(arr); i++ {

	}

	return 0, false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // инициализируем отсортированный массив.
	target := 7                                 // элемент для поиска.

	// выполняем бинарный поиск.
	index := binarySearch(arr, target)

	// выводим результат поиска.
	if index != -1 {
		fmt.Printf("Element %d found at index %d.\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array.\n", target)
	}
}

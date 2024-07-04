package main

import (
	"fmt"
)

// binarySearch выполняет бинарный поиск элемента target в отсортированном массиве arr
// возвращает индекс и булевое значение с результатом поиска
func binarySearch(arr []int, target int) (int, bool) {
	lenght := len(arr)

	if lenght == 0 {
		return 0, false
	}

	middle := lenght / 2

	current := middle

	for {
		curVal := arr[current]
		if target > curVal {
			current++
			if current >= lenght {
				return 0, false
			}
			if target < arr[current] {
				return 0, false
			}
		} else if target < curVal {
			current--
			if current < 0 {
				return 0, false
			}
			if target > arr[current] {
				return 0, false
			}
		} else {
			return current, true
		}
	}

	return 0, false
}

func main() {
	arr := []int{1, 7} // инициализируем отсортированный массив.
	target := 1        // элемент для поиска.

	// выполняем бинарный поиск.
	index, found := binarySearch(arr, target)

	// выводим результат поиска.
	if found {
		fmt.Printf("Element %d found at index %d.\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array.\n", target)
	}
}

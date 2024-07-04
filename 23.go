package main

import "fmt"

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	// Пример слайса
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента для удаления
	i := 2

	// Удаление элемента
	newSlice := remove(slice, i)

	// Вывод результата
	fmt.Println("Original slice:", slice)
	fmt.Println("New slice after removing element at index", i, ":", newSlice)

}

package main

import "fmt"

func main() {
	// Исходная последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество, используя карту
	set := make(map[string]struct{})

	// Добавляем элементы в множество
	for _, item := range sequence {
		set[item] = struct{}{}
	}

	// Выводим множество
	fmt.Println("Set:", set)
}

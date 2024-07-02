package main

import (
	"fmt"
)

// Функция для пересечения двух множеств
func intersection(set1, set2 []int) []int {
	// Создаем карту для хранения элементов первого множества
	setMap := make(map[int]struct{})
	for _, v := range set1 {
		setMap[v] = struct{}{}
	}

	// Создаем срез для хранения пересечения
	var result []int
	for _, v := range set2 {
		if _, found := setMap[v]; found {
			result = append(result, v)
			// Удаляем элемент из карты, чтобы избежать дублирования в результате
			delete(setMap, v)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersection(set1, set2)
	fmt.Println("Intersection:", result)
}

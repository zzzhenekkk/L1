package main

import (
	"fmt"
)

// partition разделяет массив на две части относительно опорного элемента (pivot)
// и возвращает индекс опорного элемента после разделения.
func partition(arr []int, low, high int) int {
	pivot := arr[high] // выбираем опорный элемент (pivot) в конце массива.
	i := low - 1       // инициализируем индекс меньшего элемента.

	for j := low; j < high; j++ {
		// если текущий элемент меньше опорного.
		if arr[j] < pivot {
			i++                             // увеличиваем индекс меньшего элемента.
			arr[i], arr[j] = arr[j], arr[i] // меняем местами элементы arr[i] и arr[j].
		}
	}
	// меняем местами опорный элемент с элементом после последнего меньшего элемента.
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // возвращаем индекс опорного элемента.
}

// iterativeQuickSort выполняет нерекурсивную быструю сортировку с использованием стека.
func iterativeQuickSort(arr []int) {
	stack := make([]int, len(arr)) // создаем стек для хранения границ подмассивов.

	top := -1 // инициализируем верхушку стека.

	// добавляем начальные границы всего массива в стек.
	top++
	stack[top] = 0 // левая граница.
	top++
	stack[top] = len(arr) - 1 // правая граница.

	// пока стек не пуст.
	for top >= 0 {
		high := stack[top] // получаем правую границу из стека.
		top--
		low := stack[top] // получаем левую границу из стека.
		top--

		// разделяем массив на две части и получаем индекс опорного элемента.
		p := partition(arr, low, high)

		// если левая часть не пустая, добавляем ее границы в стек.
		if p-1 > low {
			top++
			stack[top] = low
			top++
			stack[top] = p - 1
		}

		// если правая часть не пустая, добавляем ее границы в стек.
		if p+1 < high {
			top++
			stack[top] = p + 1
			top++
			stack[top] = high
		}
	}
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Before sorting:", arr)

	iterativeQuickSort(arr)
	fmt.Println("After sorting: ", arr)
}

//// быстрая сортировка но с рекурсией, будет переполнен стек вызовов если больше размеры в

// partition разделяет массив на две части относительно опорного элемента (pivot)
// и возвращает индекс опорного элемента после разделения.
func partition2(arr []int, low, high int) int {
	pivot := arr[high] // выбираем опорный элемент (pivot) в конце массива.
	i := low - 1       // инициализируем индекс меньшего элемента.

	for j := low; j < high; j++ {
		// если текущий элемент меньше опорного.
		if arr[j] < pivot {
			i++                             // увеличиваем индекс меньшего элемента.
			arr[i], arr[j] = arr[j], arr[i] // меняем местами элементы arr[i] и arr[j].
		}
	}
	// меняем местами опорный элемент с элементом после последнего меньшего элемента.
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // возвращаем индекс опорного элемента.
}

// quickSort выполняет рекурсивную быструю сортировку.
func quickSort2(arr []int, low, high int) {
	if low < high {
		// разделяем массив на две части и получаем индекс опорного элемента.
		p := partition2(arr, low, high)

		// рекурсивно сортируем левую часть массива.
		quickSort2(arr, low, p-1)
		// рекурсивно сортируем правую часть массива.
		quickSort2(arr, p+1, high)
	}
}

func main2() {
	arr := []int{3, 6, 8, 10, 1, 2, 1} // инициализируем массив.
	fmt.Println("Before sorting:", arr)

	quickSort(arr, 0, len(arr)-1)       // выполняем быструю сортировку.
	fmt.Println("After sorting: ", arr) // выводим отсортированный массив.
}

//В начале создается массив для сортировки и инициализируется стек для хранения границ подмассивов.
//Индекс top используется для отслеживания верхушки стека.
//В стек добавляются начальные границы всего массива (от 0 до len(arr)-1).
//Основной цикл сортировки:
//
//Пока стек не пуст, из него извлекаются границы текущего подмассива (low и high).
//Вызов функции partition делит текущий подмассив на две части относительно опорного элемента (pivot).
//Возвращаемое значение p - это индекс, по которому pivot помещается на свое место в отсортированном массиве.
//После разделения, если в подмассиве есть элементы слева от pivot, эти границы добавляются в стек.
//Если есть элементы справа от pivot, их границы также добавляются в стек.
//Процесс повторяется до тех пор, пока все подмассивы не будут отсортированы.
//Функция partition:
//
//Эта функция отвечает за выбор опорного элемента и разбиение массива на две части: элементы, меньшие pivot, и элементы,
//большие или равные pivot.
//Элементы массива перемещаются так, чтобы все элементы меньшие pivot были слева, а большие или равные pivot - справа.
//После завершения перемещения, pivot ставится на свое правильное место в отсортированном массиве, и функция возвращает
//его индекс.
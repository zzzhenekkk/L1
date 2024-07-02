package main

import (
	"fmt"
)

// Установка i-го бита в 1
func setBit(n int64, i uint) int64 {
	return n | (1 << i)
}

// Установка i-го бита в 0
func clearBit(n int64, i uint) int64 {
	return n &^ (1 << i)
}

func main() {
	var num int64 = 0 // Начальное значение

	// Устанавливаем 3-й бит в 1
	num = setBit(num, 3)
	fmt.Printf("После установки 3-го бита в 1: %064b\n", num)

	// Устанавливаем 5-й бит в 1
	num = setBit(num, 5)
	fmt.Printf("После установки 5-го бита в 1: %064b\n", num)

	// Очищаем 3-й бит
	num = clearBit(num, 3)
	fmt.Printf("После установки 3-го бита в 0: %064b\n", num)

	// Очищаем 5-й бит
	num = clearBit(num, 5)
	fmt.Printf("После установки 5-го бита в 0: %064b\n", num)
}

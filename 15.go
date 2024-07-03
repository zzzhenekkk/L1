package main

import (
	"fmt"
	"unicode/utf8"
)

//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

//Данный фрагмент кода может привести к нескольким негативным последствиям:
//
//Проблема с выделением памяти:
//
//Переменная justString хранит подстроку из createHugeString, которая ссылается на ту же область памяти, что и оригинальная
//строка v. Это значит, что вся память, выделенная для v, не будет освобождена, пока justString используется.
//Возможная проблема с некорректным использованием Unicode:
//
//Строки в Go являются UTF-8, и простое обрезание строки может нарушить кодировки символов, если обрезка происходит
//посередине многобайтового символа.

var justString string

func createHugeString(size int) string {
	// Эта функция создает строку заданного размера.
	// Для примера, будем использовать строку, состоящую из повторяющихся символов "A".
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1 << 10) // Создаем большую строку размером 1024 байта
	if utf8.ValidString(v) {
		runes := []rune(v) // Преобразуем строку в срез рун, чтобы работать с символами Unicode
		if len(runes) > 100 {
			justString = string(runes[:100]) // Берем первые 100 рун и создаем новую строку
		} else {
			justString = v // Если строка короче 100 рун, берем всю строку
		}
	} else {
		// Обработка невалидной UTF-8 строки, если это необходимо
		justString = ""
	}
}

func main() {
	someFunc()
	fmt.Println(justString) // Выводим первые 100 символов строки

	asd := "asdasdadasdasd"
	fmt.Println(asd)
	runes := []rune(asd)

	fmt.Printf("%c", runes[0])

}

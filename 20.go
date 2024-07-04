package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func reverseWords(str string) string {
	arrStr := strings.Fields(str)
	lenght := len(arrStr)
	reversedArrStr := make([]string, lenght)
	for i, value := range arrStr {
		reversedArrStr[lenght-1-i] = value
	}
	return strings.Join(reversedArrStr, " ")
}

func main() {
	str := reverseWords("snow dog sun")
	fmt.Println(str)
}

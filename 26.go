package main

import (
	"fmt"
	"strings"
)

func UniqCharacters(str string) bool {
	lowStr := strings.ToLower(str)
	lowStrRunes := []rune(lowStr)
	mRunes := make(map[rune]struct{}, len(lowStrRunes))
	for _, char := range lowStrRunes {
		_, exists := mRunes[char]
		if exists {
			return false
		} else {
			mRunes[char] = struct{}{}
		}
	}
	return true
}

func main() {

	str := "abcd"
	fmt.Println("Уникальные ли буквы в строке ", str, " :", UniqCharacters(str))
}

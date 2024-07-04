package main

import "fmt"

func flipLine(str string) string {

	runes := []rune(str)

	lenght := len(runes)

	reversed := make([]rune, lenght)

	for i, r := range runes {
		reversed[lenght-1-i] = r
	}

	return string(reversed)
}

func main() {
	str := flipLine("главрыба")
	fmt.Println(str)
}

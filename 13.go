package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	i := 0
	j := 1

	i, j = j, i

	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
}

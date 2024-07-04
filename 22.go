package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1048577", 10)
	b.SetString("2097153", 10)

	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма: %s + %s = %s\n", a.String(), b.String(), sum.String())

	diff := new(big.Int).Sub(a, b)
	fmt.Printf("Разность: %s - %s = %s\n", a.String(), b.String(), diff.String())

	product := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение: %s * %s = %s\n", a.String(), b.String(), product.String())

	quotient := new(big.Int).Div(a, b)
	fmt.Printf("Частное: %s / %s = %s\n", a.String(), b.String(), quotient.String())

}

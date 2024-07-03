package main

import (
	"fmt"
	"reflect"
)

// через type switch
func identifyType1(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("The type of the variable is int and the value is %d\n", v)
	case string:
		fmt.Printf("The type of the variable is string and the value is %s\n", v)
	case bool:
		fmt.Printf("The type of the variable is bool and the value is %t\n", v)
	case chan int:
		fmt.Printf("The type of the variable is channel of int\n")
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main1() {
	var a int = 42
	var b string = "hello"
	var c bool = true
	var d chan int = make(chan int)

	identifyType1(a)
	identifyType1(b)
	identifyType1(c)
	identifyType1(d)
	identifyType1(3.14) // Пример неизвестного типа
}

// через рефлекторы

func identifyType(i interface{}) {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Int:
		fmt.Printf("The type of the variable is int and the value is %d\n", i)
	case reflect.String:
		fmt.Printf("The type of the variable is string and the value is %s\n", i)
	case reflect.Bool:
		fmt.Printf("The type of the variable is bool and the value is %t\n", i)
	case reflect.Chan:
		fmt.Printf("The type of the variable is channel\n")
	default:
		fmt.Printf("Unknown type: %s\n", t.Kind())
	}
}

func main2() {
	var a int = 42
	var b string = "hello"
	var c bool = true
	var d chan int = make(chan int)

	identifyType(a)
	identifyType(b)
	identifyType(c)
	identifyType(d)
	identifyType(3.14) // Пример неизвестного типа
}

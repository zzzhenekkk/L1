package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action
//от родительской структуры Human (аналог наследования).

type Human struct {
	name string
	age  int
}

func (h *Human) say() {
	fmt.Println("Hello, i'm ", h.name)
}

type Cat struct {
	name string
}

func (c *Cat) say() {
	fmt.Println("Meow!")
}

type Action struct {
	Human
	//Cat
}

func (a *Action) act() {
	a.Human.say()
	//a.Cat.say()
	println(a.name)
}

func main() {
	actor := Action{
		Human{"Evgeniy", 28},
		//Cat{"Ryzhik"},
	}
	actor.act()
	actor.say()

}

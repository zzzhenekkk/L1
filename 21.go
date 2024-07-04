package main

import "fmt"

// Target - это интерфейс, который ожидает клиент
type Target interface {
	Request() string
}

// Adaptee - это существующий класс с несовместимым интерфейсом
type Adaptee struct{}

// SpecificRequest - метод Adaptee, который нужно адаптировать
func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter - адаптер, который делает интерфейс Adaptee совместимым с интерфейсом Target
type Adapter struct {
	*Adaptee
}

// Request - метод, который адаптирует SpecificRequest к интерфейсу Target
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	// Создаем объект Adaptee
	adaptee := &Adaptee{}

	// Создаем адаптер, который делает Adaptee совместимым с Target
	adapter := &Adapter{Adaptee: adaptee}

	// Клиент использует интерфейс Target, но на самом деле работает с Adaptee через Adapter
	fmt.Println(adapter.Request())
}

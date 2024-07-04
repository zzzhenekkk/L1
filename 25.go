package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	start := time.Now()
	for time.Since(start) < d {
	}
}

func main() {
	fmt.Println("Начало работы")
	Sleep(2 * time.Second)
	fmt.Println("Прошло 2 секунды")
}

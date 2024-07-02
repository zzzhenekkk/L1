package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины

// ////////// 1  Использование канала для сигнала о завершении работы:
func worker1(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Goroutine stopped")
			return
		default:
			fmt.Printf("Working...")
			time.Sleep(time.Second)
		}
	}
}

func main1() {
	stopCh := make(chan struct{})
	go worker1(stopCh)

	time.Sleep(4 * time.Second)
	close(stopCh)
	time.Sleep(2 * time.Second)
}

/////////// 2  Использование контекста (context.Context):

func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped")
			return
		default:
			fmt.Println("working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main2() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)

}

/////////// 3   Использование тикера:

func worker3() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	tickCount := 0

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("working...")
			tickCount++
			fmt.Println("Tick at", t)
			if tickCount >= 5 {
				fmt.Println("Stopping goroutine after 5 ticks")
				return
			}
		}
	}
}

func main3() {
	go worker3()
	time.Sleep(5 * time.Second)

}

/////////// 4   Использование таймера:

func main4() {
	// Завершающий канал для остановки горутины
	done := make(chan struct{})

	// Горутина, которая завершится через 2 секунды
	go func() {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Goroutine finished after 2 seconds")
			close(done)
		}
	}()

	// Основная функция ждет сигнала завершения от горутины
	<-done
	fmt.Println("Main function exiting")
}

//////////// 5 Использование переменной-флага Atomic:

var stopFlag int32

func worker5() {
	for {
		if atomic.LoadInt32(&stopFlag) == 1 {
			fmt.Println("Goroutine stopped")
			return
		}
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main5() {
	go worker5()

	time.Sleep(2 * time.Second)
	atomic.StoreInt32(&stopFlag, 1)
	time.Sleep(1 * time.Second)
}

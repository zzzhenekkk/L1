package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
//которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

func processWorker(countWorkers int, data <-chan interface{}, done <-chan struct{}) {

	var wg sync.WaitGroup

	for i := 1; i <= countWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case value, ok := <-data:
					if !ok {
						return
					}
					fmt.Println("Worker:", i, "Value:", value)
				case <-done:
					return
				}

			}
		}()
	}

	wg.Wait()

}

func main() {
	var numWorkers int
	fmt.Print("Enter number of workers:")
	fmt.Scanf("%d", &numWorkers)

	chanData := make(chan interface{})
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				chanData <- i
				i++
			}
		}

	}()

	go func() {
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				str := fmt.Sprintf("строка %d", i)
				chanData <- str
				i++
				//time.Sleep(time.Second)
			}
		}

	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	defer close(sigChan)

	go func() {
		<-sigChan
		fmt.Println("Received SIGINT, shutting down...")
		close(done)

	}()
	processWorker(numWorkers, chanData, done)

	close(chanData)
	fmt.Println("Program terminated")

}

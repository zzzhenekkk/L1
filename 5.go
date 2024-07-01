package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("Enter number of seconds to run the program: ")
	fmt.Scanf("%d", &n)

	chanI := make(chan int)
	//chanSignal := make(chan os.Signal)
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case <-done:
				close(chanI)
				return
			default:
				chanI <- i
				i++
				time.Sleep(time.Second)
			}

		}
	}()

	go func() {
		for value := range chanI {
			fmt.Println(value)
		}
	}()

	time.AfterFunc(time.Duration(n)*time.Second, func() {
		fmt.Println("Time's up!")
		close(done)
	})

	//signal.Notify(chanSignal, syscall.SIGINT)
	<-done
	//<-chanSignal
	//close(done)

	//signal.Stop(chanSignal)
	//close(chanSignal)

}
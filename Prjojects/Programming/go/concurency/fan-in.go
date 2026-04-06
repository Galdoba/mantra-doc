package main

import (
	"fmt"
	"sync"
)

func SplitChannels[T any](inputChls <-chan T, n int) []<-chan T {
	outputChls := make([]chan T, n)
	for i := 0; i < n; i++ {
		outputChls[i] = make(chan T)
	}

	go func() {
		idx := 0
		for value := range inputChls {
			outputChls[idx] <- value //может быть не блокирующим
			idx = (idx + 1) % n

		}

		for _, ch := range outputChls {
			close(ch)
		}

	}()

	//can't cast []chan T to []<-chan T
	resultChl := make([]<-chan T, n)
	for i := 0; i < n; i++ {
		resultChl[i] = outputChls[i]
	}

	return resultChl
}

func main() {
	channel := make(chan int)

	go func() {
		defer func() {
			close(channel)
		}()

		for i := 0; i < 100; i += 3 {
			channel <- i
		}
	}()

	channels := SplitChannels(channel, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for value := range channels[0] {
			fmt.Println("ch1:", value)
		}
	}()

	go func() {
		defer wg.Done()
		for value := range channels[1] {
			fmt.Println("ch2:", value)
		}
	}()

	wg.Wait()

}

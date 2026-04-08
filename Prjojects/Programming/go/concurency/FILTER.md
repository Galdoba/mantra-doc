---
updated_at: 2026-04-06T20:58:19.000+10:00
---
аналог TRANSFORMER

```go
package main

import (
	"fmt"
)

func Filter[T any](inputChs <-chan T, conditionFunc func(T) bool) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)
		for number := range inputChs {
			if conditionFunc(number) {
				outputCh <- number
			}
		}
	}()

	return outputCh
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)

		for i := 0; i < 100; i += 3 {
			channel <- i
		}
	}()

	isOdd := func(value int) bool {
		return value%2 != 0
	}

	for number := range Filter(channel, isOdd) {
		fmt.Println(number)
	}

}
```
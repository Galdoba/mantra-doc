---
updated_at: 2026-04-06T21:13:41.658+10:00
---
```go

package main

import (
	"fmt"
	"sync"
	"time"
)

// generate – создаёт канал и отправляет в него значения (как и раньше)
func generate[T any](values ...T) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)
		for _, value := range values {
			outputCh <- value
		}
	}()

	return outputCh
}

// processParallel запускает n воркеров, обрабатывающих данные параллельно
func processParallel[T any](inputCh <-chan T, actionFunc func(T) T, workersCount int) <-chan T {
	outputCh := make(chan T)

	var wg sync.WaitGroup

	// Запускаем workersCount горутин-воркеров
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		workerID := i + 1 // для отладки

		go func(id int) {
			defer wg.Done()
			for value := range inputCh {
				// Имитируем долгую операцию (например, обращение к БД или API)
				time.Sleep(100 * time.Millisecond)
				result := actionFunc(value)
				fmt.Printf("Воркер %d обработал %v -> %v\n", id, value, result)
				outputCh <- result
			}
		}(workerID)
	}

	// Закрываем выходной канал, когда все воркеры закончат
	go func() {
		wg.Wait()
		close(outputCh)
	}()

	return outputCh
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mul := func(value int) int {
		return value * value
	}

	fmt.Println("Запуск с 3 параллельными воркерами (добавлена задержка 100 мс):")
	start := time.Now()

	for result := range processParallel(generate(values...), mul, 3) {
		fmt.Println("Результат из main:", result)
	}

	elapsed := time.Since(start)
	fmt.Printf("Общее время выполнения: %v\n", elapsed)
}
```
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	nums := make([]int, size)

	for i := range size {
		nums[i] = rand.Int()
	}

	return nums
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	result := data[0]

	for _, v := range data {
		if v > result {
			result = v
		}
	}

	return result
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) <= CHUNKS {
		return maximum(data)
	}

	maxFromChunks := make([]int, CHUNKS)
	size := len(data) / CHUNKS

	var wg sync.WaitGroup

	for i := range CHUNKS {
		start := i * size
		end := start + size

		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			maxFromChunks[i] = maximum(data[start:end])
		}()
	}

	wg.Wait()

	return maximum(maxFromChunks)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	now := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(now).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	now = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(now).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}

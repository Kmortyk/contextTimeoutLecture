package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func main() {
	size := 100000000
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(10)
	}

	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan int)
	go sumArrayAndPrint(ctx, resultChan, &array)

	// обрабатываем Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	select {
	case <-sigChan:
		fmt.Println("Ctrl + C catch")
		cancel()
	case res := <-resultChan:
		fmt.Printf("result sum: %v\n", res)
	}
}

func sumArrayAndPrint(ctx context.Context, resultChan chan int, array *[]int) {
	length := len(*array)
	partsNum := 100

	step := length / partsNum
	results := make([]int, partsNum)

	for idx := 0; idx < partsNum; idx++ {
		start := idx * step
		end := start + step

		if idx == partsNum-1 {
			end += length % partsNum
		}

		part := (*array)[start:end]

		go func(idx int) {
			results[idx] = sum(ctx, &part)
		}(idx)
	}

	resultChan <- sum(ctx, &results)
}

func sum(ctx context.Context, array *[]int) int {
	sum := 0
	for i := 0; i < len(*array); i++ {
		// на каждой итерации
		select {
		case <-ctx.Done():
			return 0
		default:
		}
		sum += (*array)[i]
	}
	return sum
}

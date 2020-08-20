package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func main() {
	size := 10_000_000
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Int()
	}

	ctx, cancel := context.WithCancel(context.Background())
	resultChan := sumParallel(ctx, &array)
	sigChan := make(chan<- os.Signal)

	// обрабатываем Ctrl+C
	signal.Notify(sigChan, os.Interrupt)

	select {
	case <- sigChan:
		cancel()
	case res := <- resultChan:
		fmt.Printf("result sum: %v\n", res)
	}
}

func sumArrayAndPrint(ctx context.Context, array *[]int) {
	length := len(*array)
	partsNum := 100

	step := length / partsNum
	results := make([]int, partsNum)

	for idx := 0; idx < partsNum; idx++ {
		start := idx * step
		end := start + step

		if idx == partsNum - 1 {
			end += length % partsNum
		}

		part := (*array)[start:end]

		go func(idx int) {
			results[idx] = sum(ctx, &part)
		}(idx)
	}

	fmt.Println(sum(ctx, &results))
}

func sum(ctx context.Context, array *[]int) int {
	sum := 0
	for i := 0; i < len(*array); i++ {
		// на каждой итерации
		if _, ok := <- ctx.Done(); ok {
			return 0
		}
		sum += (*array)[i]
	}
	return sum
}
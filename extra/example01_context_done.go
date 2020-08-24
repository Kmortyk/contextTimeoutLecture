package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

/*
	Можно перехватывать сигналы операционной системы и
	завершать программу в штатном режиме.
*/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan:
			fmt.Println("\nCatch CTRL + C !")
			cancel()
		case <-ctx.Done():
		}
	}()

	const size = 100_000_000
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(10)
	}
	sum := sumArrayAndPrint(ctx, array)
	fmt.Println(sum)
}

func sumArrayAndPrint(ctx context.Context, array []int) int {
	length := len(array)
	const partsNum = 100

	step := length / partsNum
	results := make([]int, partsNum)

	res := make(chan struct{})

	for idx := 0; idx < partsNum; idx++ {
		start := idx * step
		end := start + step

		if idx == partsNum-1 {
			end += length % partsNum
		}

		part := array[start:end]

		go func(idx int) {
			results[idx] = sum(ctx, &part)
			res <- struct{}{}
		}(idx)
	}

	i := 0
	for range res {
		i++
		if i >= partsNum {
			break
		}
	}

	return sum(ctx, &results)
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

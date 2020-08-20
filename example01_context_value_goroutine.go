package main

import (
	"context"
	"fmt"
	"time"
)

/*
	context.WithValue позвоялет передавать значения на несколько уровней ниже
	вместе с контекстом.

	Все `функциональные` контексты образуются из базового, путём оборачивания
	в соответствующие методы.
 */

func main() {
	workers := 5
	ctx := context.WithValue(context.Background(), "workerNum", workers)

	for workerIdx := 0; workerIdx < workers; workerIdx++ {
		go func(ctx context.Context, idx int) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Printf("working [%v/%v]\n", idx + 1, ctx.Value("workerNum"))
					time.Sleep(time.Second)
				}
			}
		}(ctx, workerIdx)
	}

	time.Sleep(time.Second * 5)
}
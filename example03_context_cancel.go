package main

import (
	"context"
	"fmt"
	"time"
)

/*
	context.WithCancel возвращает функцию, с помощью которой
	можно сообщить дочерним процессам, что им необходимо
	завершить свою работу.

	В интерфейсе Context отсутствует метод Cancel, отмену может
	произвести только метод, который создал контекст.
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	endCh := make(chan struct{}, 3)

	go Worker(ctx, "1", endCh)
	go Worker(ctx, "2", endCh)
	go Worker(ctx, "3", endCh)

	// Working time
	time.Sleep(time.Second * 5)

	cancel()

	_, _, _ = <-endCh, <-endCh, <-endCh
}

func Worker(ctx context.Context, data string, endCh chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker properly ended up.")
			endCh <- struct{}{}
			return
		default:
			time.Sleep(time.Second * 1)
			fmt.Println(data)
		}
	}
}

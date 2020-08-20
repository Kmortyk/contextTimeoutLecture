package main

import (
	"context"
	"fmt"
	"time"
)

/*
	При добавлении новой обёртки WithCancel, каждый новый ребёнок добавляется
	в мапу, чтобы при отмене родителя, все дочерние контексты также бросили завершение.
 */

func main() {
	ctx1, cancel := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx1)

	go func() {
		<- ctx1.Done()
		fmt.Println("Worker 1 cancelled")
	}()

	go func() {
		<- ctx2.Done()
		fmt.Println("Worker 2 cancelled")
	}()

	cancel()

	time.Sleep(time.Millisecond)
}
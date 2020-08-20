package main

import (
	"context"
	"fmt"
	"time"
)

/*
	context.WithDeadline - позволяет установить временной порог, зайдя за который,
	в ctx.Done() придёт сигнал об завершении работы.

	context.WithTimeout - обёртка над WithTimeout.
*/

func main() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1 * time.Second))
	go func(){
		<- ctx.Done()
		fmt.Println("Worker 1 ended up.")
	}()

	ctx, _ = context.WithTimeout(context.Background(), 1 * time.Second)
	go func(){
		<- ctx.Done()
		fmt.Println("Worker 2 ended up.")
	}()

	time.Sleep(time.Second * 5)
}
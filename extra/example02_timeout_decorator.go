package main

import (
	"fmt"
	"time"
)

/*
	Можно воспользоваться паттерном decorator для оборачивания функций
	в timeout.
*/

func handleFunction(message string) string {
	// делаем некоторые тяжёлые операции
	time.Sleep(4 * time.Second)

	return fmt.Sprintf("hello, %v", message)
}

func timeoutDecorator(handler func(string) string) {
	timeout := time.After(3 * time.Second)
	resultCh := make(chan string)

	go func() {
		resultCh <- handler("world!")
	}()

	select {
	case <-resultCh:
		fmt.Println("job done")
	case <-timeout:
		panic("timeout")
	}
}

func main() {
	timeoutDecorator(handleFunction)
}

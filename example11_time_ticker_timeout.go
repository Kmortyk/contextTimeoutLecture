package main

import (
	"fmt"
	"time"
)

/*
	time.NewTicker - создаёт новый `тикер`, лучше пользоваться time.Tick.
	С помощью select'а можно легко объединить tick и timeout.
*/

func main() {
	// создаём переменые вне цикла !!!
	done := make(chan bool)

	timeout := time.After(1 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		time.Sleep(1 * time.Second)
		done <- true
	}()

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf(".")
		case <-timeout:
			fmt.Println("timeout!!!")
			return
		case <-done:
			fmt.Println("job is done")
			return
		}
	}
}

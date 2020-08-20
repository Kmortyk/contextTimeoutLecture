package main

import (
	"fmt"
	"time"
)

/*
	time.Sleep приостанавливает поток выполнения на некоторое время.
*/

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("woke up!")
	}()

	time.Sleep(2 * time.Second)
}

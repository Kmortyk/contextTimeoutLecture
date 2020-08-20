package main

import (
	"fmt"
	"time"
)

/*
	time.After создаёт канал, в который приходит сигнал после определённого
	промежутка времени.
*/

func main() {
	c := make(chan string, 1)

	select {
	case m := <-c:
		fmt.Printf("We got message %v\n", m)
	case <-time.After(3 * time.Second):
		fmt.Println("timed out")
	}
}

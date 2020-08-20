package main

import (
	"fmt"
	"time"
)

/*
	time.Tick создаёт новый объект Ticker и возвращает канал из него.
*/

func main() {
	tick := time.Tick(1 * time.Second)

	for next := range tick {
		fmt.Printf("update at: %v\n", next.Format("15:04:05"))
	}
}

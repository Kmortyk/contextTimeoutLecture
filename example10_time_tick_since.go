package main

import (
	"fmt"
	"time"
)

/*
	time.Tick срабатывает после отсчёта первого временного интервала.
*/

func main() {
	tick := time.NewTicker(500 * time.Millisecond)
	defer tick.Stop()

	t := time.Now()

	for range tick.C {
		fmt.Println("in loop", time.Since(t))
	}
}

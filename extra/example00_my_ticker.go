package main

import (
	"fmt"
	"time"
)

func newTicker() <- chan bool {
	ticker := make(chan bool, 1)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			ticker <- true
		}
	}()
	return ticker
}

func main()  {
	ticker := newTicker()

	for {
		select {
		case <-ticker:
			fmt.Println("tick")
		}
	}
}
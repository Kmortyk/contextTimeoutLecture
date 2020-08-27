package main

import (
	"fmt"
	"time"
)

/*
	time.AfterFunc выполняет некоторую функцию через указанное время.
	Функция запускается в отдельной горутине.
*/

func main() {
	t := time.Now()
	end := make(chan struct{})

	time.AfterFunc(3*time.Second, func() {
		fmt.Printf("passed %v\n", time.Since(t))
		end <- struct{}{}
	})

	<-end
}

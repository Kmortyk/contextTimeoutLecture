package main

import (
	"context"
	"fmt"
)

/*
	context.WithValue оборачивает переданный контекст в структуру
	с одной парой key - value.
*/

func main() {
	ctx1 := context.WithValue(context.Background(), "a", "1")
	fmt.Println(ctx1)

	ctx2 := context.WithValue(ctx1, "b", "2")
	fmt.Println(ctx2)

	ctx3 := context.WithValue(ctx2, "c", "3")
	fmt.Println(ctx3)

	///
	ctx4 := context.WithValue(ctx2, "d", "4")
	fmt.Println(ctx4)

	///
	fmt.Println(ctx4.Value("f"))
}
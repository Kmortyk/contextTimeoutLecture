package main

import (
	"context"
	"encoding/json"
	"fmt"
)

/*
	context.Background - это просто число с методами context.

	То же самое можно сказать и об context.TODO, его использут,
	на этапе проектировапния, кога непонятно, какой из контекстов
	будет использоваться в будущем.

	48 ASCII char - '0'
 */

func main() {
	ctx := context.Background()
	fmt.Println(ctx)
	fmt.Println(json.Marshal(ctx))

	ctxTodo := context.TODO()
	fmt.Println(ctxTodo)
	fmt.Println(json.Marshal(ctxTodo))
}
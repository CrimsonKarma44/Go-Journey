package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("doSomething", ctx.Value("myKeys"))

	anotherCtx := context.WithValue(ctx, "myKeys", "anotherValue")
	doAnother(anotherCtx)

	fmt.Println("doSomething: myKeys' value is", ctx.Value("myKeys"))
}

func doAnother(ctx context.Context) {
	fmt.Println("doAnother", ctx.Value("myKeys"))
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "myKeys", "myValue")
	doSomething(ctx)
	fmt.Println(ctx.Done())
}

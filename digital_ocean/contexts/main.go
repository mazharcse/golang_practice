package main

import (	
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "mykey", "myvalue")
	fmt.Println("hello context")
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Println("Do something:", ctx.Value("mykey"))
}
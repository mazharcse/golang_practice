package main

import (	
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()
	fmt.Println("hello context")
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Println("Do something:", ctx)
}
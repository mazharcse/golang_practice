package main

import (
	"fmt"
)

const p = 100

var a = 10

func outer() func() {
	money := 100
	age := 40
	fmt.Println("Money, Age: ", money, age)

	show := func() {
		money = money + a + p
		fmt.Println("Money: ", money)
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}


func main() {
	fmt.Println("Closure practice")

	call()
}
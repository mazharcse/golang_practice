package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func multipy(a int, b int) int {
	return a * b
}

func processOperation(a int, b int, operation func(int, int) int) int {
	result := operation(a, b)

	return result
}

func returnFunction() func(int, int) int {
	return add
}

func main() {
	fmt.Println("Higher oder function practice")

	var number1 int
	var number2 int	
	fmt.Println("Enter first number: ")
	fmt.Scanln(&number1)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&number2)

	sum := processOperation(number1,number2, add)

	multiplication := processOperation(number1, number2, multipy)

	fmt.Println("summation: ", sum)
	fmt.Println("multiplication: ", multiplication)


	addition := returnFunction()
	result := addition(number1,number2)
	fmt.Println("addition returns add func: ", result)
}
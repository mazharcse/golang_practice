package main

import (
	"fmt"
	"package_scope/mathlib"
)

func main() {
	fmt.Println("Welcome to package scope testing")

	var number1 int
	var number2 int	
	fmt.Println("Enter first number: ")
	fmt.Scanln(&number1)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&number2)

	
	sum := mathlib.Add(number1,number2)

	fmt.Println(sum)
}
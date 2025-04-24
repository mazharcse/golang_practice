package main

import (
	"fmt"
	"runtime"
)

type Person struct {
	Name string
	Age	 uint8
}

func createPerson(name string, age uint8) *Person {
	// allocate
	return &Person{Name:name, Age:age}
}

func main() {
	fmt.Println("Dyanamic memory allocation")

	pNumber := new(int) // allocated on the heap
    *pNumber = 42

	fmt.Println(*pNumber)

	pNumber = nil


	p := createPerson("Bob", 30)

	fmt.Println(p.Name, p.Age)

	p = nil
	// Force GC for demonstration (do not do this in produciton)
    runtime.GC()
}
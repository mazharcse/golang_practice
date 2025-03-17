package main

import (
	"fmt"
)

// base struct
type Animal struct {
	Name string
}

// method for animal
func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

// Cat embeds Aminal composition
type Cat struct {
	Animal
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "Meawo")
}

// Dog embeds Aminal composition
type Dog struct {
	Animal
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "Barks")
}


func main() {
	cat := Cat{Animal{Name: "Tushi"}}
	cat.Speak()

	dog := Dog{Animal{Name: "Tommmy"}}
	dog.Speak()
}
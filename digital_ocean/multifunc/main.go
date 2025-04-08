package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	total := 100
	go printNumbers(total, &wg)
	go generateNumbers(total, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
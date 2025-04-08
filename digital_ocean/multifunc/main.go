package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
		ch <- idx
	}
}

func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// for idx := 1; idx <= total; idx++ {
	// 	fmt.Printf("Printing number %d\n", idx)
	// }
	for num := range ch {
		fmt.Printf("read %d from channel\n", num)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	total := 100

	intChan := make(chan int)

	go printNumbers(intChan, &wg)
	generateNumbers(total, intChan, &wg)

	//wg.Wait()
	
	close(intChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}

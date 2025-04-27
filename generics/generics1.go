package main

import "fmt"

// Generic function to find the maximum of two values
func Max[T int | float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Max(10, 20))       // 20 (int)
    fmt.Println(Max(3.14, 2.71))   // 3.14 (float64)
}
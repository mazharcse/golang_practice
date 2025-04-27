package main

import (
    "fmt"
    "golang.org/x/exp/constraints"
)

// // Ordered is a constraint that permits any ordered type: any type
// // that supports the operators < <= >= >.
// // If future releases of Go add new ordered types,
// // this constraint will be modified to include them.
// //
// // Note that floating-point types may contain NaN ("not-a-number") values.
// // An operator such as == or < will always report false when
// // comparing a NaN value with any other value, NaN or not.
// // See the [Compare] function for a consistent way to compare NaN values.
// type Ordered interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64 |
// 		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
// 		~float32 | ~float64 |
// 		~string
// }



// Constraint for ordered types (supports <, >, etc.)
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Min(10, 20))       // 10 (int)
    fmt.Println(Min(3.14, 2.71))   // 2.71 (float64)
    fmt.Println(Min("a", "b"))     // a (string)
}


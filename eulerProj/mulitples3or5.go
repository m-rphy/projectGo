package main 

import "fmt"

// If we list all the natural numbers below 10 
// that are multiples of 3 or 5, we get 3,5,6 and 9. 
// The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
    n := 1000

    result := 0

    for i := 0; i < n; i++ {
        if i%3 == 0 || i%5 == 0 {
            result += i
            //fmt.Println(result)
        }
    }

    fmt.Println(result)
}

package main

import "fmt"

func main() {
    n := 10
    var ints []int

    for i := 0; i < n; i++ {
        ints = append(ints, i)
    }

    for _, val := range ints {
        if val % 2 == 0 {
            fmt.Printf("%v is even \n", val)
        } else {
            fmt.Printf("%v is odd \n", val)
        }
    }

}

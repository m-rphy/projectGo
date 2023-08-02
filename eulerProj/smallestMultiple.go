package main

import "fmt"

// 2520 is the smallest number that can be 
// divided by each of the numbers from  to  without any remainder.

// What is the smallest positive number that 
// is evenly divisible by all of the numbers from  to ?

func smallestMultiple() int {
    num := 2520
    smallest := false

    for smallest != true {
        num++
        smallest = divisibleBy20(num)
    }

    return num
}

func divisibleBy20(num int) bool {


    for i := 1; i <= 20; i++ {
        if num%i != 0 {
            return false
        }
    }
    return true
}

func main() {
    result := smallestMultiple()
    fmt.Printf("The Smallest number divisible by every number \n between 1 and 20 is %d", result)
}

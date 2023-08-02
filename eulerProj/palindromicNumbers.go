package main

import (
	"fmt"
	"strconv"
)

// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers
// is 9009 = 91 * 99.

// Find the largest palindrome made from the product of two 3-digit numbers


func palindromicNumbers() int {

    largest := 0

    for i := 999; i >= 100; i-- {
        for j := 999; j > 100; j-- {
            temp := i * j
            valid := palindromChecker(temp)
            if valid == true && temp > largest {
                largest = temp
            }
        }
    }
    return largest 
}

func palindromChecker(num int) bool {
    
    // Use the strconv method on the the variable to 
    // convert it to string and store the result in 
    // a seperate variable.

    str := strconv.Itoa(num)
    
    l, r := 0, len(str) - 1

    for l <= r {
        if str[l] != str[r] {
            return false
        } else {
            l++
            r--
        }
    }
    return true
}

func main() {
    
    result := palindromicNumbers()

    fmt.Printf("The largest palindromic number made from 3-digit numbers is %d", result)
}

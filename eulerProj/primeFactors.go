package main

import "fmt"

// The prime factors of 12195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143?

func primeFactorizor() []int {

    prime := 600851475143
    
    i := 2
    var result []int

    for i*i < prime {
        if prime%i == 0 {
            result = append(result, i)
            prime /= i
        } else {
            i++
        }
    }

    if prime > 1 {
        result = append(result, prime)
    }

    return result
}

func main() {
//	var num int
//	fmt.Print("Enter a number: ")
//	fmt.Scan(&num)
    
    prime := 600851475143
	primes := primeFactorizor()
	fmt.Printf("The prime factor(s) of %d are: %v\n", prime, primes)
    fmt.Printf("And the largest is %d", primes[len(primes)-1])
}

package main

import "fmt"


// Each new term in the Fibonacci sequence 
// is generated by adding the previous two 
// terms. By starting with  and , the first terms will be:
//
// By considering the terms in the Fibonacci sequence 
// whose values do not exceed four million, find the 
// sum of the even-valued terms.


// This function calculates the nthFibonacci number
// through memoized recursion
func nthFibonacci(num int, cache map[int]int) (int, map[int]int)  {
    // Initialize the cache with the initial values
	if len(cache) == 0 {
		cache[0] = 1
		cache[1] = 1
	}

	// break case 1;
	// if the number passed is less than or equal to 1, fill cache with num
	if num <= 1 {
		cache[num] = num
		return num, cache
	}

	// break case 2:
	// if num passed is in the cache, return cached value
	if val, ok := cache[num]; ok {
		return val, cache
	}

	// recursive case:
	// else, fill cache with newly computed value;
    left, cache := nthFibonacci(num-1, cache) 
    right, cache :=  nthFibonacci(num-2, cache)

    cache[num] = left + right

	return cache[num], cache
}

func findFibonacciSeriesLengeth() (int,int) {
    n := 4000000
    cache := make(map[int]int)

    l := 0

    el := 0

    for el < n {
        el, cache = nthFibonacci(l, cache)
        l++
    }

    return l, el
}

func sumFibEven() int {
    
    result := 0
    cache := make(map[int]int)
    
    for i := 0; i < 35; i++ {
        temp, updatedCache := nthFibonacci(i, cache)
        cache = updatedCache
        if temp%2 == 0 {
            result += temp
        }
    }

    return result
}

func main() {
    sumResult := sumFibEven()
	result, el := findFibonacciSeriesLengeth()
    fmt.Printf("The length of the series is %v, with a value of %v \n", result, el)
    fmt.Println(sumResult)
}

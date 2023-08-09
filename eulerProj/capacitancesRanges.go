package main

import "fmt"


// Function to generate all possible combinations of capacitors
func generateCombinations(n int, arr []int, i int, currComb []int) {

    // basecase:
    if i == n {
        *arr = append(*arr, currComb)
        return 
    }
    // 0 is for series
    currComb = append(currComb, 0)
    generateCombinations(n, , i+1, currComb)

    // 1 is for parallel
    currComb = append(currComb, 1)
    generateCombinations(n, arr, i+1, currComb)

}

func processComb(arr []int, c float64) float64 {
    
    temp_s := 0.0
    temp_p := 0.0

    for _, val := range arr {
        
        if val == 0 {
            temp_s += 1.0/c
        } else {
            temp_p += c
        }
    }

    total := 1.0/ temp_s + temp_p
    return total
}


func main() {

    n := 3
    var arr []int

    comb := generateCombinations(n, arr, 0)
    result := processComb(comb, 60.0)

    fmt.Println(result)
}

package main

import "fmt"

func processComb(comb []int, cap float64) float64 {
    total := 0.0

    for _, val := range comb {
        if val == 1 {
            total += cap
        } else {
            total += 1/cap
        }
    }
    return total
}

// Function to generate all possible combinations of capacitors
func generateCombinations(n int, arr []int, i int) []int {

    comb := arr
    // basecase:
    if i == n {
        return comb
    }
    
    comb[i] = 0
    comb = generateCombinations(n, comb, i+1)


    comb[i] = 1
    comb = generateCombinations(n, comb, i+1)
    
    return comb
}

func main() {
	// Define the capacitance of each individual capacitor
	cap := 60.0 // Replace this with the capacitance value in Farads

	n := 3 // Replace this with the number of identical capacitors

    var arr []int
    comb := generateCombinations(n, arr, 0)

    for i := 0; i<=n; i++ {
        c := processComb(comb, cap)
        fmt.Print(c)

    }

}

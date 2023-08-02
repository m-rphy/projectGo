package main

import "fmt"


func inParallel(cap []float64) float64 {
    total := 0.0

    for _, c := range cap {
        total += 1/c
    }
    return 1/ total
}

func inSeries(cap []float64) float64 {

    total := 0.0

    for _, c := range cap {
        total += c
    }
    return total
}

func generateCombinations(cap []float64, n int) [][]float64 {
    
    if n == 0 {
        return [][]float64{{}}
    }
    
    if len(cap) == 0 {
		return nil
	}

	first, rest := cap[0], cap[1:]

	var combinations [][]float64

	for _, sub := range generateCombinations(rest, n-1) {
		combinations = append(combinations, append([]float64{first}, sub...))
	}
    fmt.Print(combinations)
	combinations = append(combinations, generateCombinations(rest, n)...)

	return combinations

}

func main() {
	// Define the capacitance of each individual capacitor
	capacitance := 60.0 // Replace this with the capacitance value in Farads

	n := 3 // Replace this with the number of identical capacitors

	// Generate all possible combinations of capacitors
	capacitances := make([]float64, n)
	for i := 0; i < n; i++ {
		capacitances[i] = capacitance
	}
    
	allCombinations := generateCombinations(capacitances, n)

	// Calculate and print all possible capacitances
	fmt.Printf("All possible capacitances of %d identical capacitors with capacitance %f Farads:\n", n, capacitance)
	for _, comb := range allCombinations {
		seriesCap := inSeries(comb)
		parallelCap := inParallel(comb)
		fmt.Printf("Series: %f Farads, Parallel: %f Farads\n", seriesCap, parallelCap)
        fmt.Println(allCombinations)
	}
}

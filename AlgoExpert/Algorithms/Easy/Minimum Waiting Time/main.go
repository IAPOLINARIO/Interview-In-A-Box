package main

import "fmt"

func main() {

	queries := []int{3, 2, 1, 2, 6}
	result := MinimumWaitingTime(queries)

	fmt.Printf("Minimum waiting time: %d \n", result)
}

//MinimumWaitingTime calculates the minimum waiting time
func MinimumWaitingTime(queries []int) int {

	if len(queries) == 1 {
		return 0
	}

	orderedQueries := sortAsc(queries)

	fmt.Printf("Query: %d, Waiting Time: %d \n", orderedQueries[0], 0)

	waitingTime := orderedQueries[0]
	result := 0
	for i := 1; i < len(orderedQueries); i++ {
		fmt.Printf("Query: %d, Waiting Time: %d \n", orderedQueries[i], waitingTime)
		result += waitingTime
		waitingTime += orderedQueries[i]
	}

	return result
}

func sortAsc(values []int) []int {
	ordered := false
	for !ordered {
		ordered = true
		for i := 1; i < len(values); i++ {
			if values[i] < values[i-1] {
				values[i], values[i-1] = values[i-1], values[i]
				ordered = false
			}
		}
	}

	return values
}

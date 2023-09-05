package main

import "fmt"

func main() {

	inputChanges := []int{5, 7, 1, 1, 2, 3, 22} //20
	//inputChanges := []int{1, 1, 1, 1, 1} //6

	fmt.Print(NonConstructibleChange(inputChanges))
}

//O(n log n) time
//O(1) space

// NonConstructibleChange returns a non-constructible change
func NonConstructibleChange(coins []int) int {

	if len(coins) == 0 {
		return 1
	}

	sortedCoins := sortAsc(coins)

	if sortedCoins[0] > 1 {
		return 1
	}

	sumChanges := 0
	for i := 0; i < len(sortedCoins)-1; i++ {
		sumChanges += sortedCoins[i]

		if sortedCoins[i+1] > (sumChanges + 1) {
			return sumChanges + 1
		}
	}
	sumChanges += sortedCoins[len(sortedCoins)-1]

	return sumChanges + 1
}

func sortAsc(items []int) []int {
	allSorted := false

	for !allSorted {
		allSorted = true
		for i := 0; i < len(items)-1; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				allSorted = false
			}
		}
	}
	return items
}

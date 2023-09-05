package main

import "fmt"

func main() {

	inputArray := []int{5, 1, 22, 25, 6, -1, 8, 10}
	inputSequence := []int{1, 6, -1, -1}

	fmt.Print(IsValidSubsequence(inputArray, inputSequence))
}

// IsValidSubsequence validates the sequence
func IsValidSubsequence(array []int, sequence []int) bool {
	checkedSequence := 0
	if len(sequence) <= len(array) {
		lastIndex := 0
		for i := 0; i < len(sequence); i++ {
			for x := lastIndex; x < len(array); x++ {
				if sequence[i] == array[x] {
					lastIndex = x + 1
					checkedSequence = checkedSequence + 1
					break
				}
			}
		}
	}

	return checkedSequence == len(sequence)
}

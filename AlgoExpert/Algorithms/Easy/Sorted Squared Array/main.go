package main

import (
	"fmt"
	"math"
)

// O(n) time | O(n) space

func main() {

	input := []int{1, 2, 3, 5, 6, 8, 9} //[1 4 9 25 36 64 81]
	//input := []int{-2, -1}             //[1, 4]
	//input := []int{-5, -4, -3, -2, -1} //[1, 4, 9, 16, 25]

	fmt.Print(SortedSquaredArray(input))

}

// SortedSquaredArray sort an array
func SortedSquaredArray(array []int) []int {
	result := make([]int, len(array))

	firstIndex := 0
	lastIndex := len(array) - 1
	currentIndex := len(array) - 1

	for currentIndex >= 0 {
		if math.Abs(float64(array[lastIndex])) > math.Abs(float64(array[firstIndex])) {
			result[currentIndex] = array[lastIndex] * array[lastIndex]
			lastIndex--
		} else {
			result[currentIndex] = array[firstIndex] * array[firstIndex]
			firstIndex++
		}

		currentIndex--
	}

	return result
}

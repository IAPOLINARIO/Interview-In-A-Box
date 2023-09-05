package main

import "fmt"

func main() {
	input := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	fmt.Print(twoSum(input, target))
}

func twoSum(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]
		for y := i + 1; y < len(nums); y++ {
			nextNumber := nums[y]

			fmt.Printf("Current sum: %v + %v = %v \n", currentNum, nextNumber, currentNum+nextNumber)
			if currentNum+nextNumber == target {
				result = append(result, currentNum)
				result = append(result, nextNumber)
				return result
			}

		}

		result = []int{}
	}

	return result

}

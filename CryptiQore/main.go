// Write a script that prints each number from 1 to 100 on a new line. For each multiple of 3, print "Coffe" instead of the number.
// For each multiple of 5, print "Break" instead of the number. For numbers which are multiples of both 3 and 5, print "CoffeBreak" instead of the number

package main

import (
	"fmt"
	"os"
)

const head = "H"
const body = "."
const tail = "T"

func main() {
	input := "H..T..H "

	result, error := checkSnake(input)

	if error != nil {
		fmt.Printf("Error: %s", error.Error())
		os.Exit(1)
	}

	if result {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}

}

func checkSnake(input string) (bool, error) {
	hasHead := false
	hasBody := false
	hasTail := false

	for _, v := range input {
		currentChar := fmt.Sprintf("%c", v)

		if hasHead && currentChar == head {
			return false, nil
		}

		if currentChar == head {
			hasHead = true
		} else if currentChar == tail {
			hasTail = true
		} else {
			hasBody = true
		}
	}

	if hasHead && hasTail || (!hasHead && !hasTail && hasBody) {
		return true, nil
	}

	return false, nil
}

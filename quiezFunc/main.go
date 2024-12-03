package main

import (
	"errors"
	"fmt"
)

func main() {
	// Calculate the total of scores
	score := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	total := sum(score)
	fmt.Println("Total:", total)

	// Test calculate function with various operators
	operators := []string{"+", "-", "*", "/", "="} // Test for all operators
	for _, operator := range operators {
		result, err := calculate(10, 2, operator) // Call the calculate function
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Result of 10 %s 2: %d\n", operator, result)
		}
	}
}

func calculate(number, number2 int, operator string) (int, error) {
	// Perform operation based on the operator
	switch operator {
	case "+":
		return number + number2, nil
	case "-":
		return number - number2, nil
	case "*":
		return number * number2, nil
	case "/":
		return number / number2, nil
	default:
		// Return an error if the operator is invalid
		return 0, errors.New("Invalid operator")
	}
}

func sum(numbers []int) int {
	// Calculate the total sum of all numbers in the slice
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

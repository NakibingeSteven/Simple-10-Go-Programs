package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//check if there are enough  command line arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator <number1> <operator> <number2>")
		return
	}

	//parse commadn line arguments
	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	// Check for parsing errors
	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing input:", err1, err2)
		return
	}

	// Perform the calculation based on the operator
	result := 0.0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		// Check for division by zero
		if num2 == 0 {
			fmt.Println("Error: Division by zero.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator. Supported operators are +, -, *, /")
		return
	}

	// Display the result
	fmt.Printf("%v %s %v = %v\n", num1, operator, num2, result)
}

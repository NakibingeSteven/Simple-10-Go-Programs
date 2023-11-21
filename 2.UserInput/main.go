package main

import "fmt"

func main() {
	//use input store varible
	var userInput string

	//prompt the user
	fmt.Println("Enter Somthing")

	//The _ (underscore) is the blank identifier in Go. It's used when you want to ignore the value returned by a function.
	//In this case, fmt.Scan returns two values: the number of items successfully scanned (n) and an error (err).
	// The blank identifier _ is used to ignore the count of items.
	//read user input
	_, err := fmt.Scan(&userInput)

	//if error is not non
	if err != nil {
		fmt.Println("Erro reading input: ", err)
		return
	}

	//dsisplay the user input
	fmt.Println("you entered: ", userInput)

}

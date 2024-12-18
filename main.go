package main

import "fmt"

func main() {
	var num int

	// Input the initial number
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Loop to keep adding 1 after each user input
	for {
		// Add 1 to the number
		num++

		// Display the updated number
		fmt.Println("Updated number:", num)

		// Prompt the user to input another number or quit
		fmt.Print("Enter another number (or type 'quit' to stop): ")
		var input string
		fmt.Scan(&input)

		// Check if the user wants to quit
		if input == "quit" {
			fmt.Println("Exiting program.")
			break
		}
	}
}

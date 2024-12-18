package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	sum := 0

	// Loop to calculate the sum of numbers from 1 to n
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Printf("The sum of numbers from 1 to %d is: %d\n", n, sum)
}

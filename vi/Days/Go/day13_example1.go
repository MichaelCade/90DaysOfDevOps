package main

import "fmt"

func main() {

	const DaysTotal int = 90
	var remainingDays uint = 90
	challenge := "#90DaysOfDevOps"

	fmt.Printf("Welcome to the %v challenge.\nThis challenge consists of %v days\n", challenge, DaysTotal)

	var TwitterName string
	var DaysCompleted uint

	// asking for user input
	fmt.Println("Enter Your Twitter Handle: ")
	fmt.Scanln(&TwitterName)

	fmt.Println("How many days have you completed?: ")
	fmt.Scanln(&DaysCompleted)

	// calculate remaining days
	remainingDays = remainingDays - DaysCompleted

	fmt.Printf("Thank you %v for taking part and completing %v days.\n", TwitterName, DaysCompleted)
	fmt.Printf("You have %v days remaining for the %v challenge\n", remainingDays, challenge)
	fmt.Println("Good luck")
}

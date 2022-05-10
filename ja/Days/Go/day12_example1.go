package main

import "fmt"

func main() {
	
	challenge := "#90DaysOfDevOps"
	const daystotal = 90

	fmt.Printf("Welcome to %v\n", challenge)
	fmt.Printf("This is a %v challenge\n", daystotal)

	var TwitterName string
	var DaysComplete int
	// ask user for their twitter handle

	TwitterName = "@MichaelCade1"
	DaysComplete = 12
	fmt.Printf("%v has completed %v days of the challenge\n", TwitterName, DaysComplete)
	fmt.Println("Great work")
}

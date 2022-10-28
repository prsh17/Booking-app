// refer GO Documentation

// import package
package main

import (
	"fmt"
	"strings"
)

// entry point is required (main function)
func main() {

	// const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	// var bookings = []string{}
	// bookings := []string{}

// Go have only for loops
// for{
// 	#Infinite loop
// }


var userFName string
	var userLName string
	var userTickets uint
	
	// input
	fmt.Print("Enter your first name: ")
	fmt.Scan(&userFName) // need pointer to take input
	
	fmt.Print("Enter your last name: ")
	fmt.Scan(&userLName)
	
	fmt.Print("Enter no. of tickets: ")
	fmt.Scan(&userTickets)
	
	// logic to book tickets
	remainingTickets -= userTickets
	bookings = append(bookings, userFName+" "+userLName)
	
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Type of slice: %T\n", bookings)
	fmt.Printf("Length of slice: %v\n", len(bookings))
	
	// for-each
	fName := []string{}
	for _, bookings := range bookings{
		var name = strings.Fields(bookings)
	    fName = append(fName, name[0])
	}
	fmt.Printf("All bookings: %v\n", fName)

}

// 1st you need to initialize the project
// for that we use go cmd --> go mod init project-name
// refer GO Documentation
// IN GO EVERYTHING IS ORGSNIZED INTO PACKAGES

// import package
package main

import (
	"fmt"
	"strings"
)

// entry point is required (main function)
func ma() {
	var conferenceName string = "Go Conference"
	// conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	// fmt.Printf("%T\n",conferenceName) // type of var

	// call function
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var userFName string
		var userLName string
		var userEmail string
		var userTickets uint

		// input
		fmt.Print("Enter your first name: ")
		fmt.Scan(&userFName) // need pointer to take input

		fmt.Print("Enter your last name: ")
		fmt.Scan(&userLName)

		fmt.Print("Enter your Email: ")
		fmt.Scan(&userEmail)

		fmt.Print("Enter no. of tickets: ")
		fmt.Scan(&userTickets)

		// input validation
		isValidName := len(userFName) >= 2 && len(userLName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketsNumber {

			// logic to book tickets
			remainingTickets -= userTickets
			bookings = append(bookings, userFName+" "+userLName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", userFName, userLName, userTickets, userEmail)
			fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)

			// print first names with func
			fname := getFirstNames(bookings)
			fmt.Printf("All bookings: %v\n", fname)

			if remainingTickets == 0 {
				// end pr0garm
				fmt.Printf("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address is invalid.")
			}
			if !isValidTicketsNumber {
				fmt.Println("No. of tickets you entered is invaild.")
			}
		}

	}
	// Go have switch caes
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {

	fmt.Printf("Welcome to %v booking application!!\n", confName)
	fmt.Println("We have total of", confTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames(bookings []string) []string {
	fName := []string{}
	for _, bookings := range bookings {
		var name = strings.Fields(bookings)
		fName = append(fName, name[0])
	}
	return fName
}

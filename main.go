// 1st you need to initialize the project
// for that we use go cmd --> go mod init project-name
// refer GO Documentation
// IN GO EVERYTHING IS ORGSNIZED INTO PACKAGES

// import package
package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

// package lvl var
var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50

// var bookings []string //slice type
// var bookings = make([]map[string]string, 0) //map type // list of map thats y we have [] //need to initialise it
var bookings = make([]UserData, 0) //struct 

// map datatype
// firstname: ""
// lastname: ""
// email: ""
// tickets: 3

// structure --for mix datatype,with properties
type UserData struct{
	firstname string
	lastname string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// entry point is required (main function)
func main() {

	// fmt.Printf("%T\n",conferenceName) // type of var

	// call function
	greetUsers()

	// for { 
		// input
		firstname, lastname, email, userTickets := getUserInput()
		// input validation
		isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstname, lastname, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketsNumber {

			// logic to book tickets
			bookTicket(userTickets, firstname, lastname, email)

			wg.Add(1)
			// send tickets
			go sendTickets(userTickets, firstname, lastname, email)
			// go keyword 

			// print first names with func
			fname := getFirstNames()
			fmt.Printf("All bookings: %v\n", fname)

			if remainingTickets == 0 {
				// end progarm
				fmt.Printf("Our conference is booked out. Come back next year.")
				// break
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

	// }
	// Go have switch case
	wg.Wait()
}

func greetUsers() {

	fmt.Printf("Welcome to %v booking application!!\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	fName := []string{}
	for _, bookings := range bookings {
		// fName = append(fName, bookings["firstname"])// for map
		fName = append(fName, bookings.firstname)// for struct
	}
	return fName
}

func getUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTickets uint

	// input
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstname) // need pointer to take input

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Print("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter no. of tickets: ")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func bookTicket(userTickets uint, firstname string, lastname string, email string) {
	remainingTickets -= userTickets

	// map for a user
	// var userData = make(map[string]string) //map . we can't mix datatypes
	var userData = UserData{
		firstname: firstname,
		lastname: lastname,
		email: email,
		numberOfTickets: userTickets,
	}
	// userData["firstname"] = firstname // for map
	// userData["lastname"] = lastname
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // coverts str to base 10 (decimal)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)

}

func sendTickets(userTickets uint, firstname string, lastname string, email string){
	time.Sleep(30 * time.Second)// stops func for 50 sec
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastname)
	fmt.Println("-----------------------")
	fmt.Printf("Sending tickets:\n %v\n to email address %v\n", ticket, email)
	fmt.Println("-----------------------")

	wg.Done()
}

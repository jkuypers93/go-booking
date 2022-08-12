package main

import (
	"booking-app-go/shared"
	"fmt"
	"time"
	"sync"
)

// package level variables
const conferenceTickets float32 = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // creating empty list of maps - need to provide initial size (extends automatically)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	
	// get user input
	firstName, lastName, email, userTickets := getUserInput()

	// validate user input
	isValidName, isValidEmail, isValidTicketNumber := shared.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1) // threads that the main thread should wait for
		go sendTicket(userTickets, firstName, lastName, email)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation e-mail at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %s\n", remainingTickets, conferenceName)
		
		// return array of first names of attendees
		firstNames := returnFirstNames()
		fmt.Printf("These are all the first names of bookings: %v\n", firstNames)

		var noTicketsRemaining bool = remainingTickets == 0 // no need for this variable, but useful to see how it is written
		if noTicketsRemaining {
			fmt.Println("Our conference in booked out. Come back next year")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name of last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("E-mail address you entered does not contain @ sign.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickers you entered is invalid")
		}
	}
	wg.Wait()
}


func greetUsers() {
	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)

	fmt.Println("Get your tickets here to attend")

	fmt.Printf("There are %v out of %v tickets availabe!\n", remainingTickets, conferenceTickets)
}

func returnFirstNames() []string {
	firstNames := []string{}
	for _, booking := range  bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your e-mail address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	bookings = append(bookings, userData)

	fmt.Printf("New User: %v\n", userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation e-mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %s\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}
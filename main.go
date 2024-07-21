package main

import (
	helper "booking_app/common"
	"fmt"
	"sync"
	"time"
)

// Package level
// package level vars cannot be created using shorthand syntax
const ConferenceTickets = 50   // cont decleration
var conferenceName = "Go Conf" //var decleration, datatype can not be changed
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // creating empty list of maps and initializing it
var firstName string
var lastName string
var email string

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser() // using var and const

	//Slices
	// dynamic array

	for remainingTickets >= 0 && len(bookings) <= 50 {
		firstName, lastName, email, userTickets := getUserInputs()

		//Aray
		// var bookings [50]string
		// bookings[0] = firstName + " " + lastName
		// fmt.Printf("whole Array %v \n", bookings)
		// fmt.Printf("No. of bookings%v", len(bookings)) // length of array

		// calling custom module--> helper.ValidateInputs
		isValidName, isValidEmail, isValidTicketNum := helper.ValidateInputs(firstName, lastName, email, userTickets, remainingTickets)
		// handling invalid user input
		if isValidName && isValidEmail && isValidTicketNum {
			bookingTicket(userTickets, firstName, lastName)

			wg.Add(1) // waits for the launched Goroutine to finish
			// Add sets the no. of GoRoutines to wait for
			// wait Blocks until the wait group counter is 0

			go sendTicket(userTickets, firstName, lastName, email) // go keyword makes the call concurent

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conf is Sold out")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Your input Name is too short")
			} else if !isValidEmail {
				fmt.Println("Your input Email is")
			} else if !isValidTicketNum {
				fmt.Printf("We only have %v, so you can not book %v tickets \n", remainingTickets, userTickets)
			}
		}

	}
	wg.Wait() // wait for the GoRoutine to complete

}

func greetUser() {
	fmt.Println("Welcome to", conferenceName, "booking App")
	fmt.Printf("we have total of %v ,tickets and %v  are still remaining \n", ConferenceTickets, remainingTickets) // string interpolation
	fmt.Println("Get your tickets here to attend")                                                                 // printing a string
	// %v- default format
	//%s - string
	//%d - integer
}

func getUserInputs() (string, string, string, uint) {

	fmt.Println("Start New Booking")
	// ask user for there name
	fmt.Println("Provide First name")
	fmt.Scan(&firstName) // taking user input in terminal
	// & is a pointer

	fmt.Println("Provide Last name")
	fmt.Scan(&lastName) // taking user input in terminal
	// & is a pointer
	fmt.Println(&lastName) // printing thg the val of pointer where in memory variable is stored

	fmt.Println("Provide email")
	fmt.Scan(&email) // taking user input in terminal

	var userTickets uint
	fmt.Println("Enter no. of tickets")
	fmt.Scan(&userTickets)
	fmt.Printf("you have booked %v ,userTickets  for  %v of data type %T for %v   \n", userTickets, firstName, userTickets, conferenceName) // string interpolation

	return firstName, lastName, email, userTickets
}

func fetchFrstName(bookings []UserData) []string {
	var firstNames = []string{}

	// Range iterates over elements for different data structures
	// for slices and arrays it gives us index and val

	// _ is a blank identifier -- to ignore the vars you have to decleare but you are not using
	//  we are telling go that there needs to a var but we will not use it

	for _, booking := range bookings {

		// strings.Fileds()
		// splits the string with white space as seprator
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames // retun val
}

func bookingTicket(userTickets uint, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets

	// create map for users
	// map is acollection of key value pairs
	// map[keyDataType]valueDataType
	// we can not have mixed dataType in map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // type conversion

	// struct
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings %v \n", fetchFrstName(bookings))

	fmt.Printf("No. of people who made bookings %v \n", len(bookings)) // length of array
	fmt.Println("Now  remainingTickets is ", remainingTickets)

}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // demo for halting the app
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("**********")
	fmt.Printf("Sending ticket: \n %v\n to email address %v \n", ticket, email)
	fmt.Println("**********")
	wg.Done() // decresing the counter of the Goroutines to wait for
}

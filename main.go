package main

import (
	"booking-app/m/v2/helper"
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for remainingTickets > 0 {

		firstname, lastname, email, userTickets := getUserInputs()

		isValidName := helper.IsValidName(firstname, lastname)

		if !isValidName {
			fmt.Println("Inputs are not valid, Please try again. ")
			continue
		}

		if userTickets <= remainingTickets {
			bookTickets(userTickets, firstname, lastname, email)

			firstnames := getFirstName()
			fmt.Printf("The first name of booking are: %v\n", firstnames)

			if remainingTickets == 0 {
				fmt.Print("Our conference is booked out.")
				break
			}

		} else {
			fmt.Printf("We only have %v tickets, You can not book %v tickets.", remainingTickets, userTickets)
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("conferenceTicket is %T, remainingTicket is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have total of %v tokcets and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstName() []string {
	firstnames := []string{}

	for _, booking := range bookings {
		var name = booking["firstName"]
		firstnames = append(firstnames, name)
	}

	return firstnames
}

func getUserInputs() (string, string, string, uint) {

	var firstname string
	var lastname string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter tikcet counts to book: ")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func bookTickets(userTickets uint, firstname string, lastname string, email string) {

	var userData = make(map[string]string)
	userData["firstName"] = firstname
	userData["lastName"] = lastname
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData)

}

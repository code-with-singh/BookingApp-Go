package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for remainingTickets > 0 {

		firstname, lastname, email, userTickets := getUserInputs()

		if userTickets <= remainingTickets {
			bookTickets(bookings, remainingTickets, userTickets, firstname, lastname, email)

			if remainingTickets == 0 {
				fmt.Print("Our conference is booked out.")
				break
			}

		} else {
			fmt.Printf("We only have %v tickets, You can not book %v tickets.", remainingTickets, userTickets)
			continue
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
		var name = strings.Fields(booking)
		var firstname = name[0]
		firstnames = append(firstnames, firstname)
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

func bookTickets(bookings []string, userTickets uint, firstname string, lastname string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstname+" "+lastname+" "+email)

	firstnames := getFirstName()
	fmt.Printf("The first name of booking are: %v\n", firstnames)
}

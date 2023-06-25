package main

import (
	"booking-app/m/v2/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstname, lastname, email, userTickets := getUserInputs()

	isValidName := helper.IsValidName(firstname, lastname)

	if !isValidName {
		fmt.Println("Inputs are not valid, Please try again. ")
	}

	if userTickets <= remainingTickets {
		bookTickets(userTickets, firstname, lastname, email)

		wg.Add(1)
		go sendTicket(userTickets, firstname, lastname, email)

		firstnames := getFirstName()
		fmt.Printf("The first name of booking are: %v\n", firstnames)

		if remainingTickets == 0 {
			fmt.Print("Our conference is booked out.")
		}

	} else {
		fmt.Printf("We only have %v tickets, You can not book %v tickets.", remainingTickets, userTickets)
	}

	wg.Wait()
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
		firstnames = append(firstnames, booking.firstName)
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

	var userData = UserData{
		firstName:       firstname,
		lastName:        lastname,
		email:           email,
		numberOfTickets: userTickets,
	}

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData)

}

func sendTicket(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	tickets := fmt.Sprintf("%v tikcets for %v %v", userTickets, firstname, lastname)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket: \n %v \n to email address %v", tickets, email)
	fmt.Println("#####################")

	wg.Done()
}

package main

import (
	"fmt"
	"strings"
)

func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	bookings := []string{}

	for remainingTickets > 0 && len(bookings) < conferenceTickets {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// asking for user input
		fmt.Println("Enter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			// book ticket in system
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The following people have booked tickets: %v\n", firstNames)

			noTickets := remainingTickets == 0

			if noTickets {
				// no tickets left
				fmt.Println("Sorry, no tickets left. Please try again later.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Invalid name. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("Invalid email. Please try again.")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number. Please try again.")
			}
			continue
		}
	}
}

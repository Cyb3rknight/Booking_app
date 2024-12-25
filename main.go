package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Define the conference name and total number of tickets
	var conferenceName string = "GopherCon 2021"
	const conferenceTickets int = 50
	var remainingTickets int

	// Read the remaining tickets from the file
	file, err := os.Open("remaining_tickets.txt")
	if err != nil {
		// If the file does not exist, set remaining tickets to the total number of tickets
		remainingTickets = conferenceTickets
	} else {
		defer file.Close()
		var remainingTicketsStr string
		// Read the remaining tickets from the file
		fmt.Fscanf(file, "%s", &remainingTicketsStr)
		remainingTickets, _ = strconv.Atoi(remainingTicketsStr)
	}

	// Define the ticket price
	var ticketPrice float64 = 100.50

	// Print welcome messages and ticket information
	fmt.Println("Welcome to our", conferenceName, "conference booking application")
	fmt.Println("Get your tickets here to attend the conference\n")
	fmt.Println("Ticket Price: ", ticketPrice, "\n")
	fmt.Println("Tickets remaining: ", remainingTickets, "\n")

	// Check if there are no remaining tickets
	if remainingTickets == 0 {
		fmt.Println("Sorry, all tickets are sold out!")
		return
	}

	// Prompt the user to enter the number of tickets they want to purchase
	var userTickets int
	fmt.Print("Enter the number of tickets you want to purchase: ")
	fmt.Scanln(&userTickets)

	// Check if the user entered a valid number
	if userTickets <= 0 {
		fmt.Println("Please enter a valid number of tickets.")
		return
	}

	// Check if the requested number of tickets is available
	if userTickets > remainingTickets {
		fmt.Printf("Sorry, we only have %d tickets remaining.\n", remainingTickets)
		return
	}

	// Update and display the remaining tickets
	totalPrice := ticketPrice * float64(userTickets)
	fmt.Printf("You have successfully purchased %d tickets.\nPrice: %.2f$\n", userTickets, totalPrice)
	remainingTickets -= userTickets
	fmt.Printf("Tickets remaining: %d\n", remainingTickets)

	// Save the updated number of remaining tickets to the file
	file, err = os.Create("remaining_tickets.txt")
	if err != nil {
		fmt.Println("Error saving remaining tickets:", err)
		return
	}
	defer file.Close()
	file.WriteString(strconv.Itoa(remainingTickets))
}

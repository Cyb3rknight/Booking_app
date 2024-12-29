package main

//bufio package is used to read the full name, which handles spaces correctly, and removes the newline character from the end of the input.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Define the conference name and total number of tickets
	var conferenceName string = "GopherCon 2021"
	const conferenceTickets int = 50
	var remainingTickets int
	var userName string

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
	fmt.Printf("We have %v tickets available for sale\n", remainingTickets)
	fmt.Println("Get your tickets here to attend the conference\n")
	fmt.Println("Ticket Price: ", ticketPrice, "\n")
	fmt.Println("Tickets remaining: ", remainingTickets, "\n")

	// Check if there are no remaining tickets
	if remainingTickets == 0 {
		fmt.Println("Sorry, all tickets are sold out!")
		return
	}

	fmt.Print("Enter your full name: ")
	// Read the full name including spaces
	reader := bufio.NewReader(os.Stdin)
	userName, _ = reader.ReadString('\n')
	userName = userName[:len(userName)-1] // Remove the newline character

	// Prompt the user to enter the number of tickets they want to purchase
	var userTickets int
	fmt.Print("Enter the number of tickets you want to purchase: ")

	// Read the user input and store it in the userTickets variable
	// fmt.Scanln reads input from the standard input until a newline character is encountered
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
	fmt.Printf("%s have successfully purchased %d tickets.\nPrice: %.2f$\n", userName, userTickets, totalPrice)
	remainingTickets -= userTickets
	fmt.Printf("Tickets remaining: %d\n", remainingTickets)

	// Save the updated number of remaining tickets to the file
	file, err = os.OpenFile("name_userTickets.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error saving name and user tickets:", err)
		return
	}
	defer file.Close()

	if userTickets == 1 {
		file.WriteString(userName + ": " + strconv.Itoa(userTickets) + " ticket\n")
	} else if userTickets > 1 {
		file.WriteString(userName + ": " + strconv.Itoa(userTickets) + " tickets\n")
	}

	file, err = os.Create("remaining_tickets.txt")
	if err != nil {
		fmt.Println("Error saving remaining tickets:", err)
		return
	}
	defer file.Close()
	file.WriteString(strconv.Itoa(remainingTickets))
}

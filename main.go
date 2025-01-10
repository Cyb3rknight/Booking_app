package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()
	myWindow := myApp.NewWindow("Conference Ticket Booking")

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

	// Create UI elements
	welcomeLabel := widget.NewLabel(fmt.Sprintf("Welcome to our %s conference booking application", conferenceName))
	ticketsLabel := widget.NewLabel(fmt.Sprintf("We have %v tickets available for sale", remainingTickets))
	priceLabel := widget.NewLabel(fmt.Sprintf("Ticket Price: %.2f$", ticketPrice))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter your full name")

	ticketCountEntry := widget.NewEntry()
	ticketCountEntry.SetPlaceHolder("Enter number of tickets")

	// Button to handle ticket purchase
	purchaseButton := widget.NewButton("Purchase Tickets", func() {
		userName = nameEntry.Text
		userTicketsStr := ticketCountEntry.Text
		userTickets, err := strconv.Atoi(userTicketsStr)

		if err != nil || userTickets <= 0 {
			dialog.ShowInformation("Invalid Input", "Please enter a valid number of tickets.", myWindow)
			return
		}

		if userTickets > remainingTickets {
			dialog.ShowInformation("Tickets Unavailable", fmt.Sprintf("Sorry, we only have %d tickets remaining.", remainingTickets), myWindow)
			return
		}

		// Update and display the remaining tickets
		totalPrice := ticketPrice * float64(userTickets)
		remainingTickets -= userTickets

		// Show success message
		dialog.ShowInformation("Purchase Successful", fmt.Sprintf("%s has successfully purchased %d tickets.\nPrice: %.2f$\nTickets remaining: %d", userName, userTickets, totalPrice, remainingTickets), myWindow)

		// Update user tickets file
		updateUserTickets(userName, userTickets)
		// Update remaining tickets file
		updateRemainingTickets(remainingTickets)

		// Update UI labels
		ticketsLabel.SetText(fmt.Sprintf("We have %v tickets available for sale", remainingTickets))
	})

	// Layout the UI
	content := container.NewVBox(
		welcomeLabel,
		ticketsLabel,
		priceLabel,
		nameEntry,
		ticketCountEntry,
		purchaseButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func updateUserTickets(userName string, userTickets int) {
	// Read the existing user tickets file
	file, err := os.Open("name_userTickets.txt")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading user tickets file:", err)
		return
	}
	defer file.Close()

	var lines []string
	userExists := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, userName+":") {
			// Update the user's ticket count
			parts := strings.Split(line, ": ")
			existingTickets, _ := strconv.Atoi(strings.Split(parts[1], " ")[0])
			newTickets := existingTickets + userTickets
			line = fmt.Sprintf("%s: %d tickets", userName, newTickets)
			userExists = true
		}
		lines = append(lines, line)
	}

	if !userExists {
		lines = append(lines, fmt.Sprintf("%s: %d tickets", userName, userTickets))
	}

	// Save the updated user tickets to the file
	file, err = os.Create("name_userTickets.txt")
	if err != nil {
		fmt.Println("Error saving user tickets:", err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		file.WriteString(line + "\n")
	}
}

func updateRemainingTickets(remainingTickets int) {
	// Save the updated number of remaining tickets to the file
	file, err := os.Create("remaining_tickets.txt")
	if err != nil {
		fmt.Println("Error saving remaining tickets:", err)
		return
	}
	defer file.Close()
	file.WriteString(strconv.Itoa(remainingTickets))
}

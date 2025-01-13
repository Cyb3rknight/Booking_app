package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// CustomTheme defines a custom theme with grey text color
type CustomTheme struct {
	fyne.Theme
}

func (c CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameForeground {
		return color.Gray{Y: 128} // Grey color
	}
	return c.Theme.Color(name, variant)
}

func main() {
	// Create a new Fyne application
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Conference Ticket Booking")

	// Define the conference name and total number of tickets
	var conferenceName string = "GopherCon 2021"
	const conferenceTickets int = 50
	var remainingTickets int

	// Read the remaining tickets from the file
	file, err := os.Open("BOOKING_APP/remaining_tickets.txt")
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

	// Create UI elements for login
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter your email address")

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter your username")

	var loginContent, registerContent *fyne.Container

	loginButton := widget.NewButton("Login", func() {
		email := emailEntry.Text
		userName, userTickets := getUserTickets(email)

		if userName == "" {
			dialog.ShowInformation("User Not Found", "No user found with this email address. Please register first.", myWindow)
			return
		}

		// Create UI elements for ticket booking
		welcomeLabel := widget.NewLabel(fmt.Sprintf("Welcome %s to our %s conference booking application", userName, conferenceName))
		ticketsLabel := widget.NewLabel(fmt.Sprintf("You have %d tickets", userTickets))
		priceLabel := widget.NewLabel(fmt.Sprintf("Ticket Price: %.2f$", ticketPrice))
		remTickets := widget.NewLabel(fmt.Sprintf("Tickets remaining: %d", remainingTickets))
		priceTickets := widget.NewLabel("Total Price: $0.00")

		ticketCountEntry := widget.NewEntry()
		ticketCountEntry.SetPlaceHolder("Enter number of tickets")

		// Update the total price dynamically
		ticketCountEntry.OnChanged = func(content string) {
			newTickets, err := strconv.Atoi(content)
			if err == nil && newTickets > 0 {
				totalPrice := ticketPrice * float64(newTickets)
				priceTickets.SetText(fmt.Sprintf("Total Price: %.2f$", totalPrice))
			} else {
				priceTickets.SetText("Total Price: $0.00")
			}
		}

		// Button to handle ticket purchase
		purchaseButton := widget.NewButton("Purchase Tickets", func() {
			userTicketsStr := ticketCountEntry.Text
			newTickets, err := strconv.Atoi(userTicketsStr)

			if err != nil || newTickets <= 0 {
				dialog.ShowInformation("Invalid Input", "Please enter a valid number of tickets.", myWindow)
				return
			}

			if newTickets > remainingTickets {
				dialog.ShowInformation("Tickets Unavailable", fmt.Sprintf("Sorry, we only have %d tickets remaining.", remainingTickets), myWindow)
				return
			}

			// Update and display the remaining tickets
			totalPrice := ticketPrice * float64(newTickets)
			remainingTickets -= newTickets
			userTickets += newTickets

			// Show success message
			dialog.ShowInformation("Purchase Successful", fmt.Sprintf("%s has successfully purchased %d tickets.\nPrice: %.2f$\nTickets remaining: %d", userName, newTickets, totalPrice, remainingTickets), myWindow)

			// Update user tickets file
			updateUserTickets(userName, email, userTickets)
			// Update remaining tickets file
			updateRemainingTickets(remainingTickets)

			// Update UI labels
			ticketsLabel.SetText(fmt.Sprintf("You have %d tickets", userTickets))
			remTickets.SetText(fmt.Sprintf("Tickets remaining: %d", remainingTickets))
			priceTickets.SetText("Total Price: $0.00")
			ticketCountEntry.SetText("")
		})

		// Button to handle logout
		logoutButton := widget.NewButton("Logout", func() {
			myWindow.SetContent(loginContent)
		})

		// Layout the UI for ticket booking
		content := container.NewVBox(
			welcomeLabel,
			ticketsLabel,
			priceLabel,
			remTickets,
			priceTickets,
			ticketCountEntry,
			purchaseButton,
			logoutButton,
		)

		myWindow.SetContent(content)
	})

	registerButton := widget.NewButton("Register", func() {
		username := usernameEntry.Text
		email := emailEntry.Text

		if username == "" || email == "" {
			dialog.ShowInformation("Invalid Input", "Please enter both username and email address.", myWindow)
			return
		}

		// Check if the user already exists
		existingUserName, _ := getUserTickets(email)
		if existingUserName != "" {
			dialog.ShowInformation("User Exists", "A user with this email address already exists. Please log in.", myWindow)
			return
		}

		// Create a new user entry
		updateUserTickets(username, email, 0)
		dialog.ShowInformation("Registration Successful", "You have successfully registered. Please log in.", myWindow)
		myWindow.SetContent(loginContent)
	})

	// Layout the UI for login and registration options
	loginContent = container.NewVBox(
		widget.NewLabel("Conference Ticket Booking"),
		emailEntry,
		container.NewHBox(
			loginButton,
			widget.NewButton("Register", func() {
				myWindow.SetContent(registerContent)
			}),
		),
	)

	registerContent = container.NewVBox(
		widget.NewLabel("Register New Account"),
		usernameEntry,
		emailEntry,
		registerButton,
		widget.NewButton("Back to Login", func() {
			myWindow.SetContent(loginContent)
		}),
	)

	myWindow.SetContent(loginContent)
	myWindow.ShowAndRun()
}

func getUserTickets(email string) (string, int) {
	// Read the existing user tickets file
	file, err := os.Open("BOOKING_APP/name_userTickets.txt")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading user tickets file:", err)
		return "", 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) == 3 && parts[1] == email {
			userName := parts[0]
			userTickets, _ := strconv.Atoi(strings.Split(parts[2], " ")[0])
			return userName, userTickets
		}
	}

	return "", 0
}

func updateUserTickets(userName, email string, userTickets int) {
	// Read the existing user tickets file
	file, err := os.Open("BOOKING_APP/name_userTickets.txt")
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
		parts := strings.Split(line, ": ")
		if len(parts) == 3 && parts[1] == email {
			line = fmt.Sprintf("%s: %s: %d tickets", userName, email, userTickets)
			userExists = true
		}
		lines = append(lines, line)
	}

	if !userExists {
		lines = append(lines, fmt.Sprintf("%s: %s: %d tickets", userName, email, userTickets))
	}

	// Save the updated user tickets to the file
	file, err = os.Create("BOOKING_APP/name_userTickets.txt")
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
	file, err := os.Create("BOOKING_APP/remaining_tickets.txt")
	if err != nil {
		fmt.Println("Error saving remaining tickets:", err)
		return
	}
	defer file.Close()
	file.WriteString(strconv.Itoa(remainingTickets))
}

# GopherCon 2021 Ticket Booking Application

This is a simple Go application for booking tickets to the GopherCon 2021 conference. The application keeps track of the remaining tickets and allows users to purchase tickets. The remaining tickets are saved to a file so that it persists across application runs.

## First step
- Install an IDE to start writing your code (I'm using VS Code)
  https://code.visualstudio.com/download
- Install Go
  https://go.dev/doc/install

## Features

- Displays the conference name, ticket price, and remaining tickets.
- Allows users to log in with their email address.
- Allows new users to register with a username and email address.
- Displays the current ticket count for logged-in users.
- Prompts the user to enter the number of tickets they want to purchase.
- Checks if the requested number of tickets is available.
- Updates and saves the remaining tickets to a file.
- Allows users to log out and log in with another account.

## Requirements

- Go 1.16 or later

## Usage

1. Clone the repository:
    ```sh
    git clone https://github.com/Cyb3rknight/booking_app.git
    cd booking_app
    ```

2. Build the application:
    ```sh
    go build -o booking_app main.go
    ```

3. Run the application:
    ```sh
    ./booking_app
    ```

4. Follow the on-screen prompts to log in, register, and purchase tickets.

## File Structure

- [main.go](http://_vscodecontentref_/1): The main application file.
- `BOOKING_APP/remaining_tickets.txt`: A file that stores the number of remaining tickets.
- `BOOKING_APP/name_userTickets.txt`: A file that stores the user information and their ticket count.

## Example

```sh
Welcome to our GopherCon 2021 conference booking application
Get your tickets here to attend the conference

Ticket Price: 100.50$
Tickets remaining: 50

Enter your email address: user@example.com
No user found with this email address. Please register first.

Register New Account
Enter your username: John Doe
Enter your email address: user@example.com
Registration Successful. Please log in.

Enter your email address: user@example.com
Welcome John Doe to our GopherCon 2021 conference booking application
You have 0 tickets
Ticket Price: 100.50$

Enter number of tickets: 2
Purchase Successful
John Doe has successfully purchased 2 tickets.
Price: 201.00$
Tickets remaining: 48
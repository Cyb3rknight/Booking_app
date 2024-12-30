# GopherCon 2021 Ticket Booking Application

This is a simple Go application for booking tickets to the GopherCon 2021 conference. The application keeps track of the remaining tickets and allows users to purchase tickets. The remaining tickets are saved to a file so that it persist across application runs.

## First step
- Install an IDE to start writing your code (I'm using VS code)
https://code.visualstudio.com/download
- Install GO 
https://go.dev/doc/install

## Features

- Displays the conference name, ticket price, and remaining tickets.
- Prompts the user to enter the number of tickets they want to purchase.
- Checks if the requested number of tickets is available.
- Updates and saves the remaining tickets to a file.

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

4. Follow the on-screen prompts to purchase tickets.

## File Structure

- [main.go]: The main application file.
- [remaining_tickets.txt]: A file that stores the number of remaining tickets.

## Example

```sh
Welcome to our GopherCon 2021 conference booking application
Get your tickets here to attend the conference

Ticket Price: 100.50$
Tickets remaining: 50

Enter the number of tickets you want to purchase: 2
You have successfully purchased 2 tickets.
Price: 201.00$
Tickets remaining: 48

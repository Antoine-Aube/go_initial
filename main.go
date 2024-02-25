//go mod init booking-app
package main
//need to declare the execution point - need to give the go compiler a starting point
import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
													//string interpolation
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	
}